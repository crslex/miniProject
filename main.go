package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/crslex/miniProject/config"
	campaign_repo "github.com/crslex/miniProject/repository/campaign"
	campaign_service "github.com/crslex/miniProject/service/campaign"
	"github.com/gorilla/mux"
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

	// Router initialization
	router := mux.NewRouter()
	// Create server, graceful shutdown
	server := &http.Server{
		Addr:    "",
		Handler: router,
	}

	go func() {
		log.Println("Starting server...")
		if err := server.ListenAndServe(); err != nil {
			log.Fatal(err)
		}
	}()

	// wait for an interrupt signal to shut down the server gracefully
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)
	<-interrupt
	log.Println("Shutting down server...")
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatal(err)
	}
	log.Println("Server stopped.")
}
