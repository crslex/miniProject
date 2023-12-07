package model

import (
	"context"
	"time"
)

type Campaign struct {
	ID           int64     `json:"id"`
	Name         string    `json:"campaign_name"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ActivaStatus bool      `json:"active"`
}

// Campaign related services
type CampaignService interface {
	GetByID(ctx context.Context, ID int64) (*Campaign, error)
	GetByListID(ctx context.Context, ListID []int64) (*[]Campaign, error)
}

// Campaign related repository

type CampaignRepository interface {
	GetByID(ctx context.Context, ID int64) (*Campaign, error)
	GetByListID(ctx context.Context, ListID []int64) (*[]Campaign, error)
}
