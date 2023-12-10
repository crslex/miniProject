package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	model "github.com/crslex/miniProject/model/campaign"
	"github.com/nsqio/go-nsq"
)

var nsq_consumer *nsq.Consumer

func main() {

	config := nsq.NewConfig()
	nsq_consumer, _ = nsq.NewConsumer("to_redis", "ch", config)
	// Add logic to insert into redis cache
	nsq_consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body[:]))
		// If possible, convert the message that has campaign JSON
		var cachedCampaign *model.Campaign
		err := json.Unmarshal(message.Body, &cachedCampaign)
		if err != nil {
			fmt.Println("Failed to unmarshal the message from the nsq consumer")
			return err
		}

		// Send to redis
		fmt.Println("Campaign message sent to redis")
		return nil
	}))
	err := nsq_consumer.ConnectToNSQLookupd("127.0.0.1:4161")
	if err != nil {
		log.Panic("Could not connect")
	}
	// graceful stop
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	<-sigChan
	nsq_consumer.Stop()

}
