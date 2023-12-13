package campaign

import (
	"context"
	"errors"
	"log"

	model "github.com/crslex/miniProject/model/campaign"
)

type campaignService struct {
	repo model.CampaignRepository
}

func (*campaignService) GetByIDGraphQL() {
	panic("unimplemented")
}

func (*campaignService) GetByListIDGraphQL() {
	panic("unimplemented")
}

func NewCampaignService(repo model.CampaignRepository) model.CampaignService {
	// grpc.Server
	return &campaignService{
		repo: repo,
	}
}

func (c *campaignService) GetByID(ctx context.Context, ID int64) (*model.Campaign, error) {
	log.Println("GetByID service is called")
	cmp, err := c.repo.GetByID(ctx, ID)
	if err != nil {
		return nil, errors.New("Error from GetByID service " + err.Error())
	}
	return cmp, nil
}

func (c *campaignService) GetByListID(ctx context.Context, ListID []int64) ([]model.Campaign, error) {
	cmpList, err := c.repo.GetByListID(ctx, ListID)
	if err != nil {
		return nil, errors.New("Error from GetByListID service " + err.Error())
	}
	return cmpList, nil
}
