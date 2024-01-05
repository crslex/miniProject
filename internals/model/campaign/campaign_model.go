package model

import (
	"time"
)

type Campaign struct {
	ID           int64     `json:"id"`
	Name         string    `json:"campaign_name"`
	Start        time.Time `json:"start"`
	End          time.Time `json:"end"`
	ActivaStatus bool      `json:"active"`
}

type CampaignElasticSearchResponse struct {
	Index       string   `json:"_index"`
	ID          string   `json:"_id"`
	Version     int      `json:"_version"`
	SeqNo       int      `json:"_seq_no"`
	PrimaryTerm int      `json:"_primary_term"`
	Found       bool     `json:"found"`
	Source      Campaign `json:"_source"`
}

// Campaign related services

// Campaign related repository
