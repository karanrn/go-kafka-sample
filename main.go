package main

import (
	"context"

	"github.com/karanrn/go-kafka-sample/kafkaops"
)

func main() {
	ctx := context.Background()
	go kafkaops.Produce(ctx)
	kafkaops.Consume(ctx)
}
