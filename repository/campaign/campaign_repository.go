package campaign

import (
	"context"
	"fmt"

	model "github.com/crslex/miniProject/model/campaign"
	"github.com/jackc/pgx/v5/pgxpool"
)

type CampaignRepository struct {
	pgConn *pgxpool.Conn
}

func NewCampaignRepository(pgConn *pgxpool.Conn) model.CampaignRepository {
	return &CampaignRepository{
		pgConn: pgConn,
	}
}

func (c CampaignRepository) GetByID(ctx context.Context, ID int) (m *model.Campaign, e error) {
	row, err := c.pgConn.Query(ctx, "SELECT * FROM campaign WHERE id=$1", ID)

	defer row.Close()
	if err != nil {
		return nil, err
	}
	err = row.Scan(m)
	if err != nil {
		return nil, fmt.Errorf("Failure in converting to object")
	}
	return m, nil

}

func (c CampaignRepository) GetByListID(ctx context.Context, ListID []int) (m *[]model.Campaign, e error) {
	return nil, nil
}
