package main

import (
  "net"
  "log"

  "google.golang.org/grpc"
  "golang.org/x/net/context"
  pb "github.com/dreae/erato/protobuf"
)

type server struct{}

func (s *server) DoRegister(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResult, error) {
  return &pb.RegisterResult{Message: "Registered: " + in.ApiKey}, nil
}

func main() {
  lis, err := net.Listen("tcp", ":27015")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterMasterServer(s, &server{})
  s.Serve(lis)
}
