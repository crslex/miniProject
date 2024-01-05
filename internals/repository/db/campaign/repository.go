package campaign

import (
	"context"

	model "github.com/crslex/miniProject/internals/model/campaign"
	"github.com/nsqio/go-nsq"
)

type CampaignRepository interface {
	GetByID(ctx context.Context, ID int64) (*model.Campaign, error)
	GetByListID(ctx context.Context, ListID []int64) ([]model.Campaign, error)
	InitNSQConsumer(nsq_consumer *nsq.Consumer)
	GetByIDElasticSearch(ctx context.Context, ID string) (*model.Campaign, error)
	GetByListIDElasticSearch(ctx context.Context, ListID []string) ([]model.Campaign, error)
	// GetAllIDs(ctx context.Context) ([]Campaign, error)
}
