package main

import (
	"crypto/tls"
	"log"
	"net"
	"net/http"

	gr "github.com/crslex/miniProject/internals/handler/campaign/grpc"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/nsqio/go-nsq"
	"github.com/redis/go-redis/v9"

	"github.com/crslex/miniProject/config"
	c "github.com/crslex/miniProject/constant"
	hCampaign "github.com/crslex/miniProject/internals/handler/campaign"
	rCampaign "github.com/crslex/miniProject/internals/repository/db/campaign/impl"
	sCampaign "github.com/crslex/miniProject/internals/usecase/campaign/impl"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	// Create config file
	cfg, err := InitConfig(c.ConfigPath)
	if err != nil {
		log.Panic("constant reading error : ", err)
	}
	log.Println("port", cfg.Server.GRPC.Host)
	// Create db connection
	err = config.InitDB()
	if err != nil {
		log.Fatal(err)
	}
	conn := config.GetConnection()

	// Create NSQ Producer
	config := nsq.NewConfig()
	NSQProducer, _ := nsq.NewProducer("127.0.0.1:4150", config)
	// Create NSQ Consumer
	config_consumer := nsq.NewConfig()
	nsq_consumer, _ := nsq.NewConsumer(c.NSQTopic, "ch", config_consumer)

	// Create redis client
	rdb := redis.NewClient(
		&redis.Options{
			Addr:     "localhost:6379", // Update with your Redis server address
			DB:       0,                // Default DB
			Password: "",
		},
	)

	// Create elasticsearch client
	elastic_cfg := elasticsearch.Config{
		Addresses: []string{
			"https://localhost:9200",
		},
		Username: "elastic",
		Password: "s4QAc6N6JzHm6TFKu9**",
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true,
			},
		},
	}

	es_client, _ := elasticsearch.NewClient(elastic_cfg)
	// repo initialization
	campaignRepo := rCampaign.NewCampaignRepository(conn, NSQProducer, rdb, es_client)

	// Initiate NSQ Consumer
	campaignRepo.InitNSQConsumer(nsq_consumer)

	// services initialization
	campaignService := sCampaign.NewCampaignService(campaignRepo)

	// New Handler
	campaignHandler := hCampaign.NewHandler(campaignService)

	// Handler initialization
	listener, err := net.Listen("tcp", cfg.Server.GRPC.Port)
	if err != nil {
		panic(err)
	}
	log.Println("Server started", cfg.Server.GRPC.Port)
	s := grpc.NewServer()
	gr.RegisterCampaignHandlerServer(s, campaignHandler)
	reflection.Register(s)
	if err := s.Serve(listener); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

}
