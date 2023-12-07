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
	cmp, err := h.sv.GetByID(ctx, int64(req.Id))
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

func (h *GRPCHandler) GetCampaignByListID(ctx context.Context, req *gr.GetCampaignByListIDRequest) (*gr.GetCampaignByIDResponse, error) {
	log.Println("GetCampaignByID is called ! ")
	cmp, err := h.sv.GetByListID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	converted_res := &gr.GetCampaignByIDResponse{
		Campaign: []*gr.Campaign{},
	}
	for _, cmp := range *cmp {
		converted_res.Campaign = append(converted_res.Campaign, &gr.Campaign{
			Id:           cmp.ID,
			CampaignName: cmp.Name,
			Start:        timestamppb.New(cmp.Start),
			End:          timestamppb.New(cmp.End),
			Active:       cmp.ActivaStatus,
		})
	}
	return converted_res, nil
}
