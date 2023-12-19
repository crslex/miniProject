package campaign

import (
	"context"
	"log"

	gr "github.com/crslex/miniProject/handler/campaign/grpc"
	mdl "github.com/crslex/miniProject/model/campaign"
	"google.golang.org/protobuf/types/known/timestamppb"
)

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
	cmp, err := h.sv.GetByID(ctx, req.GetId())
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

func (h *gRPCHandler) GetCampaignByListID(ctx context.Context, req *gr.GetCampaignByListIDRequest) (*gr.GetCampaignByListIDResponse, error) {
	log.Println("GetCampaignByListID is called ! ")
	cmp, err := h.sv.GetByListID(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	converted_res := &gr.GetCampaignByListIDResponse{
		Campaign: []*gr.Campaign{},
	}
	for _, cmp := range cmp {
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
func (h *gRPCHandler) GetCampaignByIDElasticSearch(ctx context.Context, req *gr.GetCampaignByIDElasticSearchRequest) (*gr.Campaign, error) {
	log.Println("GetCampaignByIDElasticSearch is called ! ")

	cmp, err := h.sv.GetByIDElasticSearch(ctx, req.GetId())
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

func (h *gRPCHandler) GetCampaignByListIDElasticSearch(ctx context.Context, req *gr.GetCampaignByListIDElasticSearchRequest) (*gr.GetCampaignByListIDResponse, error) {
	log.Println("GetCampaignByListIDElasticSearch is called ! ")

	cmp, err := h.sv.GetByListIDElasticSearch(ctx, req.GetId())
	if err != nil {
		return nil, err
	}

	converted_res := &gr.GetCampaignByListIDResponse{
		Campaign: []*gr.Campaign{},
	}

	for _, cmp := range cmp {
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
