package campaign

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	model "github.com/crslex/miniProject/model/campaign"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/nsqio/go-nsq"
)

type CampaignRepository struct {
	pgConn  *pgxpool.Pool
	nsqProd *nsq.Producer
}

func NewCampaignRepository(pgConn *pgxpool.Pool, nsqProd *nsq.Producer) model.CampaignRepository {
	return &CampaignRepository{
		pgConn:  pgConn,
		nsqProd: nsqProd,
	}
}

func (c *CampaignRepository) GetByID(ctx context.Context, ID int64) (m *model.Campaign, e error) {
	// Find on redis first if found, directly send to user, else continue to 2nd phases
	// 2nd stage, search on PostgreSQL, if found
	// publish to NSQ
	connection, err := c.pgConn.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from pool ")
	}
	defer connection.Release()

	rows, err := connection.Query(ctx, "SELECT * FROM campaign WHERE id=$1", ID)
	if !rows.Next() {
		log.Println("No rows selected by the query.")
		return nil, sql.ErrNoRows
	}
	nn, err := rows.Values()
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
	err = c.nsqProd.Publish("to_redis", cmp)
	if err != nil {
		log.Println("Failed to publish to redis through nsq")
		return nil, err
	}
	// Consume
	// Return back to upper layer
	return &res, nil

}

func (c *CampaignRepository) GetByListID(ctx context.Context, ListID []int64) (m *[]model.Campaign, e error) {
	res := []model.Campaign{}
	for _, id := range ListID {
		cmp, err := c.GetByID(ctx, id)
		if err != nil {
			log.Fatal("Failed in GetByLastID using GetByID", err.Error())
		}
		res = append(res, *cmp)
	}
	return &res, nil
}
