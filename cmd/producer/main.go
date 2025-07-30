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
	topic := flag.String("topic", "", "Topic to send messages to")
	count := flag.Int("count", 10, "Number of messages to send")
	key := flag.String("key", "", "Optional key for messages")
	flag.Parse()

	if *topic == "" {
		log.Fatalf("--topic is required")
	}

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(*brokers, ","),
		Topic:   *topic,
	})
	defer writer.Close()

	for i := 0; i < *count; i++ {
		msg := kafka.Message{
			Key:   []byte(*key),
			Value: []byte(fmt.Sprintf("message %d", i)),
		}
		if err := writer.WriteMessages(context.Background(), msg); err != nil {
			log.Fatalf("failed to write message: %v", err)
		}
		fmt.Printf("sent: %s\n", msg.Value)
	}
}
