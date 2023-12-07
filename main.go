package main

import (
	"fmt"
	"log"
	"net"

	gr "github.com/crslex/miniProject/handler/campaign/grpc"

	"github.com/crslex/miniProject/config"
	campaaign_handler "github.com/crslex/miniProject/handler/campaign"
	campaign_repo "github.com/crslex/miniProject/repository/campaign"
	campaign_service "github.com/crslex/miniProject/service/campaign"
	"google.golang.org/grpc"
)

func main() {
	// Create db connection
	err := config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	conn := config.GetConnection()

	// repo initialization
	campaignRepo := campaign_repo.NewCampaignRepository(conn)

	// services initialization
	campaignService := campaign_service.NewCampaignService(campaignRepo)

	fmt.Println(campaignService)

	// Handler initialization
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	log.Println("Server started")
	s := grpc.NewServer()
	gr.RegisterCampaignHandlerServer(s, &campaaign_handler.GRPCHandler{})
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
