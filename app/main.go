package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"google.golang.org/grpc"

	emp "github.com/karanrn/go-kafka-sample/employee"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", writeMessage)
	log.Println("Serving on localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", mux))
}

func writeMessage(w http.ResponseWriter, r *http.Request) {

	if r.Method == "POST" {
		var conn *grpc.ClientConn
		var newEmp emp.EmployeeRequest
		conn, err := grpc.Dial(":50001", grpc.WithInsecure())
		if err != nil {
			log.Fatalf(err.Error())
		}

		defer conn.Close()

		e := emp.NewOperationClient(conn)

		err = json.NewDecoder(r.Body).Decode(&newEmp)
		if err != nil {
			log.Fatalf("failed to decode: %v", err)
		}

		response, err := e.WriteEmployee(context.Background(), &newEmp)
		if err != nil {
			log.Fatal(err.Error())
		}

		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, fmt.Sprintf("Employee %v loaded successfully", response.EmpId))
		return
	}

	http.Error(w, "Method not supported", http.StatusNotFound)

}
