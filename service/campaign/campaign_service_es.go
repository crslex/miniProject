package campaign

import (
	"context"
	"errors"

	model "github.com/crslex/miniProject/model/campaign"
)

func (c *campaignService) GetByIDElasticSearch(ctx context.Context, ID string) (*model.Campaign, error) {
	res, err := c.repo.GetByIDElasticSearch(ctx, ID)
	if err != nil {
		return nil, errors.New("Error from GetByID service " + err.Error())

	}

	return res, nil
}

func (c *campaignService) GetByListIDElasticSearch(ctx context.Context, ID int64) ([]model.Campaign, error) {
	panic("unimplemented")
}
