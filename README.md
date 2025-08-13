# Kafka Playground

This project provides a small Kafka cluster using Docker Compose and simple Go clients for producing and consuming messages. The cluster runs Kafka in KRaft mode without ZooKeeper.

## Prerequisites
- Docker with Compose plugin
- Go 1.24+

## Running the cluster

Start the brokers:
```bash
docker compose up -d
```
This launches three Kafka brokers listening on local ports `19092`, `19093` and `19094`.

## Interactive client

Run a single client that both consumes from and produces to a topic. Messages
you type are sent to Kafka while all received messages are printed to the
console:
```bash
go run ./cmd/client --topic test --group example
```
Use `--brokers` to override the broker list or `--key` to set a key for all
produced messages. The client will keep running until you press `Ctrl+C`.

Stop the cluster with:
```bash
docker compose down
```
