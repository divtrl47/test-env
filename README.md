# Kafka Playground

This project provides a small Kafka cluster using Docker Compose and simple Go clients for producing and consuming messages. The cluster runs Kafka in KRaft mode without ZooKeeper.

## Prerequisites
- Docker and Docker Compose
- Go 1.20+

## Running the cluster

Start the brokers:
```bash
docker-compose up -d
```
This launches three Kafka brokers listening on local ports `19092`, `19093` and `19094`.

## Producer
Send messages to a topic (optionally specifying a key to control the partition):
```bash
go run producer.go --topic test --count 5
```
You can change the broker list with `--brokers` and provide a `--key` argument.

## Consumer
Consume messages from a topic as part of a consumer group:
```bash
go run consumer.go --topic test --group example
```
Multiple consumers with the same group id will share partitions.

Stop the cluster with:
```bash
docker-compose down
```
