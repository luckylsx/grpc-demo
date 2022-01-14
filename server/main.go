package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"time"

	v1 "github.com/luckylsx/rpc-demo/server/proto/v1"

	"google.golang.org/grpc"
)

var _ v1.UserServer = &Server{}

type Server struct {
	v1.UnimplementedUserServer
	a int
}

func (s *Server) UserRegister(ctx context.Context, in *v1.UserRegisterRequest) (*v1.UserRegisterResponse, error) {
	var name string
	var age int32
	var likes []string
	name = in.GetName()
	age = in.GetAge()
	likes = in.GetLike()
	fmt.Printf("resuest name=%v, age=%v,likes=%v", name, age, likes)
	input, _ := json.Marshal(in)
	fmt.Printf("input = %v", string(input))
	resp := &v1.UserRegisterResponse{
		Id:   int64(randInt()),
		Name: "response",
		Age:  100,
	}
	return resp, nil
}

const port = ":8001"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("net listen failed, error : %v", err)
	}
	fmt.Printf("grpc listening port %v\n", port)
	s := grpc.NewServer()
	var rs = Server{}
	v1.RegisterUserServer(s, &rs)
	_ = s.Serve(lis)
}

func randInt() int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(100)
}
