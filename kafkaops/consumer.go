package kafkaops

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

// Consume reads/consumes messages from the topic
func Consume(ctx context.Context) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
		GroupID: "my-group",
	})

	for {
		msg, err := r.ReadMessage(ctx)
		if err != nil {
			panic("could not read message " + err.Error())
		}

		fmt.Println("received: ", string(msg.Value))
	}
}
