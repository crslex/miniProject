package impl

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"

	constant "github.com/crslex/miniProject/constant"
	model "github.com/crslex/miniProject/internals/model/campaign"
)

func (c *CampaignRepository) GetByListIDElasticSearch(ctx context.Context, ListID []string) ([]model.Campaign, error) {
	res := []model.Campaign{}
	for _, id := range ListID {
		cmp, err := c.GetByIDElasticSearch(ctx, id)
		if err != nil {
			log.Println("Failed in GetByListIDElasticSearch using GetByIDElasticSearch", err.Error())
			continue
		}
		res = append(res, *cmp)
	}
	return res, nil
}

func (c *CampaignRepository) GetByIDElasticSearch(ctx context.Context, ID string) (*model.Campaign, error) {
	resp, err := c.es.Get(es_index_name, ID)
	if err != nil {
		log.Println("Failed to get the documents")
		return nil, err
	}

	if !resp.IsError() {
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Println("Failed to read response body")
		}

		// var campaign *model.Campaign
		var esResponse *model.CampaignElasticSearchResponse
		if err := json.Unmarshal(bodyBytes, &esResponse); err != nil {
			log.Println("Failed to unmarshal into campaign docs")
		}
		log.Println("Got campaign from es")
		return &esResponse.Source, nil
	}

	log.Println("Continuing on postgres")

	// 2nd stage, search on PostgreSQL, if found
	connection, err := c.pgConn.Acquire(context.Background())
	if err != nil {
		log.Fatal("Error while acquiring connection from pool ")
	}
	defer connection.Release()

	res := model.Campaign{}
	err = connection.QueryRow(ctx, `SELECT id,campaign_name,"start","end", active FROM campaign WHERE id=$1`, ID).Scan(&res.ID, &res.Name,
		&res.Start, &res.End, &res.ActivaStatus)
	if err != nil {
		return nil, err
	}
	// PUBLISH TO NSQ HERE
	cmp, err := json.Marshal(res)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	err = c.nsqProd.Publish(constant.NSQTopic, cmp)
	if err != nil {
		log.Println("Failed to publish to redis and es through nsq")
		return nil, err
	}
	return &res, nil

}
