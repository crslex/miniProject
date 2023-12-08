package main

import (
	"log"
	"os"
	"os/signal"
	"syscall"

	"github.com/nsqio/go-nsq"
)

var nsq_consumer *nsq.Consumer

func main() {

	config := nsq.NewConfig()
	nsq_consumer, _ = nsq.NewConsumer("write_test", "ch", config)
	nsq_consumer.AddHandler(nsq.HandlerFunc(func(message *nsq.Message) error {
		log.Printf("Got a message: %v", string(message.Body[:]))
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
