package server

import (
  "golang.org/x/net/context"
  "github.com/dreae/erebus/lib/config"

  pb "github.com/dreae/erebus/protobuf"
)

type Server struct{}

func (s *Server) DoRegister(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResult, error) {
  return &pb.RegisterResult{Message: "Registered: " + in.ApiKey}, nil
}

func NewServer(conf *config.Config) *Server {
  return &Server{}
}
