package campaign

import (
	"context"

	model "github.com/crslex/miniProject/model/campaign"
)

func (*CampaignRepository) GetByIDElasticSearch(ctx context.Context, ID int64) (*model.Campaign, error) {
	return nil, nil
}

func (*CampaignRepository) GetByListIDElasticSearch(ctx context.Context, ID int64) ([]model.Campaign, error) {
	res := []model.Campaign{}

	return res, nil
}
