package campaign

import (
	"context"

	model "github.com/crslex/miniProject/internals/model/campaign"
)

type CampaignService interface {
	GetByID(ctx context.Context, ID int64) (*model.Campaign, error)
	GetByListID(ctx context.Context, ListID []int64) ([]model.Campaign, error)
	GetByIDGraphQL()
	GetByListIDGraphQL()
	GetByIDElasticSearch(ctx context.Context, ID string) (*model.Campaign, error)
	GetByListIDElasticSearch(ctx context.Context, ListID []string) ([]model.Campaign, error)
}
