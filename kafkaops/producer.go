package kafkaops

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"github.com/segmentio/kafka-go"
)

const (
	topic          = "new-topic"
	broker1Address = "localhost:9092"
)

// Produce writes/pushes messages to the topic
func Produce(ctx context.Context) {

	i := 0

	// Intialize the writer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	for {

		err := w.WriteMessages(ctx, kafka.Message{
			Key:   []byte(strconv.Itoa(i)),
			Value: []byte("this is message: " + strconv.Itoa(i)),
		})

		if err != nil {
			panic("could not write message " + err.Error())
		}

		fmt.Println("writes: ", i)
		i++
		time.Sleep(time.Second)
	}

}
