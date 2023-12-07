package campaign

import (
	"context"
	"fmt"
	"log"

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

func (c CampaignRepository) GetByID(ctx context.Context, ID int64) (m *model.Campaign, e error) {
	rows, err := c.pgConn.Query(ctx, "SELECT * FROM campaign WHERE id=$1", ID)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		err = rows.Scan(m)
		if err != nil {
			return nil, fmt.Errorf("Failure in converting to object")
		}
		break
	}

	return m, nil

}

func (c CampaignRepository) GetByListID(ctx context.Context, ListID []int64) (m *[]model.Campaign, e error) {
	for _, id := range ListID {
		cmp, err := c.GetByID(ctx, id)
		if err != nil {
			log.Fatal("Failed in GetByLastID using GetByID", err.Error())
		}
		*m = append(*m, *cmp)
	}
	return m, nil
}
