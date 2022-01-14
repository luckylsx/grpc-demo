package main

import (
	"context"
	"fmt"
	"log"

	v1 "github.com/luckylsx/rpc-demo/server/server/proto/v1"

	"google.golang.org/grpc"
)

const (
	address = "localhost:8001"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("grpc dial failed, err: %v\n", err)
	}
	defer conn.Close()
	userClient := v1.NewUserClient(conn)
	r, err := userClient.UserRegister(context.Background(), &v1.UserRegisterRequest{
		Name: "go-grpc",
		Age:  10,
		Like: []string{"dance"},
	})
	if err != nil {
		log.Fatalf("grpc request failed, err: %v", err)
	}
	fmt.Printf("res = %+v", r)
}
