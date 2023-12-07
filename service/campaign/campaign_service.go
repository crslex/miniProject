package campaign

import (
	"context"
	"errors"

	model "github.com/crslex/miniProject/model/campaign"
)

type CampaignService struct {
	repo model.CampaignRepository
}

func NewCampaignService(repo model.CampaignRepository) model.CampaignService {
	// grpc.Server
	return &CampaignService{
		repo: repo,
	}
}

func (c *CampaignService) GetByID(ctx context.Context, ID int64) (*model.Campaign, error) {
	cmp, err := c.repo.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.New("Error from GetByID service " + err.Error())
	}
	return cmp, nil
}

func (c *CampaignService) GetByListID(ctx context.Context, ListID []int64) (*[]model.Campaign, error) {
	cmpList, err := c.repo.GetByListID(ctx, ListID)
	if err != nil {
		return nil, errors.New("Error from GetByListID service " + err.Error())
	}
	return cmpList, nil
}
