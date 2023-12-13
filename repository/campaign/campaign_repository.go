package campaign

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"time"

	model "github.com/crslex/miniProject/model/campaign"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"
)

const (
	redis_topic_name = "to_redis"
)

type CampaignRepository struct {
	pgConn  *pgxpool.Pool
	nsqProd *nsq.Producer
	rc      *redis.Client
	es      *elasticsearch.Client
}

func NewCampaignRepository(pgConn *pgxpool.Pool, nsqProd *nsq.Producer, rc *redis.Client, es *elasticsearch.Client) model.CampaignRepository {
	return &CampaignRepository{
		pgConn:  pgConn,
		nsqProd: nsqProd,
		rc:      rc,
		es:      es,
	}
}

func (c *CampaignRepository) InitNSQConsumer(nsq_consumer *nsq.Consumer) {
	// Add logic to insert into redis cache
	nsq_consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body[:]))
		// If possible, convert the message that has campaign JSON
		var cachedCampaign *model.Campaign
		err := json.Unmarshal(message.Body, &cachedCampaign)
		if err != nil {
			fmt.Println("Failed to unmarshal the message from the nsq consumer")
			return err
		}

		// Send to redis
		err = c.rc.Set(context.Background(), fmt.Sprintf("campaign:%d", cachedCampaign.ID), message.Body, 1*time.Minute).Err()
		if err != nil {
			fmt.Println("Error caching campaign in Redis:", err)
			return nil
		}

		return nil
	}))
	err := nsq_consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
}

func (c *CampaignRepository) GetByID(ctx context.Context, ID int64) (m *model.Campaign, e error) {
	// Find on redis first if found, directly send to user, else continue to 2nd phases
	keyExists, err := c.rc.Exists(ctx, fmt.Sprintf("campaign:%d", ID)).Result()
	if err != nil {
		fmt.Println("Error checking key existence in Redis:", err)
		return nil, err
	}

	if keyExists == 1 {
		value, err := c.rc.Get(ctx, fmt.Sprintf("campaign:%d", ID)).Result()
		if err != nil {
			fmt.Println("Error getting value from redis")
			return nil, err
		}

		var retrievedCampaign model.Campaign
		err = json.Unmarshal([]byte(value), &retrievedCampaign)
		if err != nil {
			fmt.Println("Error unmarshaling JSON data:", err)
			return nil, err
		}
		fmt.Println("Retrieved campaign from redis")
		return &retrievedCampaign, nil
	}

	// 2nd stage, search on PostgreSQL, if found
	connection, err := c.pgConn.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from pool ")
	}
	defer connection.Release()

	rows, err := connection.Query(ctx, "SELECT * FROM campaign WHERE id=$1", ID)
	if err != nil {
		return nil, errors.New("failed to execute query in GetByID repository layer")
	}
	if !rows.Next() {
		log.Println("Keys not found")
		return nil, sql.ErrNoRows
	}
	nn, err := rows.Values()
	if err != nil {
		return nil, errors.New("failed to acquire values from rows acquired")
	}
	res := model.Campaign{
		ID:           int64(nn[0].(int32)),
		Name:         nn[1].(string),
		Start:        nn[2].(time.Time),
		End:          nn[3].(time.Time),
		ActivaStatus: nn[4].(bool),
	}
	// PUBLISH TO NSQ HERE
	cmp, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = c.nsqProd.Publish(redis_topic_name, cmp)
	if err != nil {
		log.Println("Failed to publish to redis through nsq")
		return nil, err
	}
	return &res, nil

}

func (c *CampaignRepository) GetByListID(ctx context.Context, ListID []int64) (m []model.Campaign, e error) {
	res := []model.Campaign{}
	for _, id := range ListID {
		cmp, err := c.GetByID(ctx, id)
		if err != nil {
			log.Println("Failed in GetByLastID using GetByID", err.Error())
			continue
		}
		res = append(res, *cmp)
	}
	return res, nil
}
