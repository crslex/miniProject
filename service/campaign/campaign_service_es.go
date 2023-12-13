package campaign

import (
	"context"

	model "github.com/crslex/miniProject/model/campaign"
)

func (*campaignService) GetByIDElasticSearch(ctx context.Context, ID int64) (*model.Campaign, error) {
	panic("unimplemented")
}

func (*campaignService) GetByListIDElasticSearch(ctx context.Context, ID int64) ([]model.Campaign, error) {
	panic("unimplemented")
}
