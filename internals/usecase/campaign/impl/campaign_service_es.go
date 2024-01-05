package impl

import (
	"context"
	"errors"

	model "github.com/crslex/miniProject/internals/model/campaign"
)

func (c *campaignService) GetByIDElasticSearch(ctx context.Context, ID string) (*model.Campaign, error) {
	res, err := c.repo.GetByIDElasticSearch(ctx, ID)
	if err != nil {
		return nil, errors.New("Error from GetByIDElasticSearch service " + err.Error())

	}

	return res, nil
}

func (c *campaignService) GetByListIDElasticSearch(ctx context.Context, ListID []string) ([]model.Campaign, error) {
	cmpList, err := c.repo.GetByListIDElasticSearch(ctx, ListID)
	if err != nil {
		return nil, errors.New("Error from GetByListIDElasticSearch service " + err.Error())
	}
	return cmpList, nil
}
