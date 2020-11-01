package main

import (
	"context"
	"encoding/json"
	"log"
	"net"
	"os"

	"github.com/segmentio/kafka-go"
	"google.golang.org/grpc"

	emp "github.com/karanrn/go-kafka-sample/employee"
)

var topic = os.Getenv("TOPIC")
var broker1Address = os.Getenv("BROKERADDRESS")
var port = os.Getenv("PORT")

type server struct {
	emp.UnimplementedOperationServer
}

// WriteEmployee writes employee request to Kafka topic
func (s *server) WriteEmployee(ctx context.Context, newEmp *emp.EmployeeRequest) (*emp.EmployeeResponse, error) {
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers: []string{broker1Address},
		Topic:   topic,
	})

	e, err := json.Marshal(newEmp)
	if err != nil {
		log.Fatalf("failed to marshal EmployeeRequest: %v", err)
		return &emp.EmployeeResponse{}, err
	}

	err = w.WriteMessages(ctx, kafka.Message{
		Key:   []byte(string(newEmp.EmpId)),
		Value: []byte(string(e)),
	})

	if err != nil {
		log.Fatalf("could not write message: %v ", err)
		return &emp.EmployeeResponse{}, err
	}

	return &emp.EmployeeResponse{EmpId: newEmp.EmpId}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to serve: %v", err)
	}

	s := grpc.NewServer()
	emp.RegisterOperationServer(s, &server{})
	if err = s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
