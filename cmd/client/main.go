package main

import (
	"bufio"
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strings"
	"sync"
	"syscall"

	"github.com/segmentio/kafka-go"
)

func main() {
	brokers := flag.String("brokers", "localhost:19092,localhost:19093,localhost:19094", "Comma-separated list of bootstrap servers")
	topic := flag.String("topic", "", "Topic to use")
	group := flag.String("group", "", "Consumer group id")
	key := flag.String("key", "", "Optional key for produced messages")
	flag.Parse()

	if *topic == "" {
		log.Fatalf("--topic is required")
	}
	if *group == "" {
		log.Fatalf("--group is required")
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigCh := make(chan os.Signal, 1)
	signal.Notify(sigCh, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigCh
		cancel()
	}()

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers: strings.Split(*brokers, ","),
		GroupID: *group,
		Topic:   *topic,
	})
	defer reader.Close()

	writer := kafka.NewWriter(kafka.WriterConfig{
		Brokers: strings.Split(*brokers, ","),
		Topic:   *topic,
	})
	defer writer.Close()

	var wg sync.WaitGroup
	wg.Add(2)

	// Consumer goroutine
	go func() {
		defer wg.Done()
		for {
			msg, err := reader.ReadMessage(ctx)
			if err != nil {
				if ctx.Err() != nil {
					return
				}
				log.Printf("failed to read message: %v", err)
				continue
			}
			fmt.Printf("%s partition=%d offset=%d key=%s value=%s\n",
				msg.Topic, msg.Partition, msg.Offset, string(msg.Key), string(msg.Value))
		}
	}()

	// Producer goroutine reading from stdin
	go func() {
		defer wg.Done()
		scanner := bufio.NewScanner(os.Stdin)
		fmt.Println("Type messages and press Enter to send. Ctrl+C to exit.")
		for scanner.Scan() {
			text := scanner.Text()
			if text == "" {
				continue
			}
			msg := kafka.Message{Key: []byte(*key), Value: []byte(text)}
			if err := writer.WriteMessages(ctx, msg); err != nil {
				log.Printf("failed to write message: %v", err)
				continue
			}
		}
		if err := scanner.Err(); err != nil && ctx.Err() == nil {
			log.Printf("stdin error: %v", err)
		}
	}()

	wg.Wait()
}
