package main

import (
	"log"

	"github.com/nsqio/go-nsq"
)

var NSQProducer *nsq.Producer

func main() {
	config := nsq.NewConfig()
	NSQProducer, _ := nsq.NewProducer("127.0.0.1:4150", config)

	err := NSQProducer.Publish("write_test", []byte("test"))
	if err != nil {
		log.Panic("Could not connect")
	}

	// graceful stop
	// sigChan := make(chan os.Signal, 1)
	// signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	// <-sigChan
	NSQProducer.Stop()
}
