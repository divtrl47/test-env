package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"strings"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokers := flag.String("brokers", "localhost:19092,localhost:19093,localhost:19094", "Comma-separated list of bootstrap servers")
	topic := flag.String("topic", "", "Topic to consume")
	group := flag.String("group", "", "Consumer group id")
	flag.Parse()

	if *topic == "" {
		log.Fatalf("--topic is required")
	}
	if *group == "" {
		log.Fatalf("--group is required")
	}

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(*brokers, ","),
		GroupID: *group,
		Topic:   *topic,
	})
	defer reader.Close()

	for {
		msg, err := reader.ReadMessage(context.Background())
		if err != nil {
			log.Fatalf("failed to read message: %v", err)
		}
		fmt.Printf("%s partition=%d offset=%d key=%s value=%s\n",
			msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
	}
}
