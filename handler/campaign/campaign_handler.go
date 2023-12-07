package campaign

import (
	"context"
	"log"

	gr "github.com/crslex/miniProject/handler/campaign/grpc"
	srv "github.com/crslex/miniProject/service/campaign"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type GRPCHandler struct {
	gr.UnimplementedCampaignHandlerServer
	sv srv.CampaignService
}

func (h *GRPCHandler) GetCampaignByID(ctx context.Context, req *gr.GetCampaignByIDRequest) (*gr.Campaign, error) {
	log.Println("GetCampaignByID is called ! ")
	cmp, err := h.sv.GetByID(ctx, int(req.Id))
	if err != nil {
		return nil, err
	}
	return &gr.Campaign{
		Id:           cmp.ID,
		CampaignName: cmp.Name,
		Start:        timestamppb.New(cmp.Start),
		End:          timestamppb.New(cmp.End),
		Active:       cmp.ActivaStatus,
	}, nil
}

func (h *GRPCHandler) GetCampaignByListID(req gr.CampaignHandler_GetCampaignByListIDServer) error {
	log.Println("GetCampaignByID is called ! ")
	// req.RecvMsg()
	// cmp, err := h.sv.GetByListID(req.)
	return nil

}
