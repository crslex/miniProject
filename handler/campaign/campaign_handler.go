package campaign

import (
	"context"
	"log"

	gr "github.com/crslex/miniProject/handler/campaign/grpc"
	mdl "github.com/crslex/miniProject/model/campaign"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// mustEmbedUnimplementedCampaignHandlerServer implements grpc.CampaignHandlerServer.

type gRPCHandler struct {
	gr.UnimplementedCampaignHandlerServer
	sv mdl.CampaignService
}

func NewHandler(srv mdl.CampaignService) gr.CampaignHandlerServer {
	return &gRPCHandler{
		UnimplementedCampaignHandlerServer: gr.UnimplementedCampaignHandlerServer{},
		sv:                                 srv,
	}
}

func (h *gRPCHandler) GetCampaignByID(ctx context.Context, req *gr.GetCampaignByIDRequest) (*gr.Campaign, error) {
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

func (h *gRPCHandler) GetCampaignByListID(ctx context.Context, req *gr.GetCampaignByListIDRequest) (*gr.GetCampaignByIDResponse, error) {
	log.Println("GetCampaignByListID is called ! ")
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
