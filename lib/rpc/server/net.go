package server

import (
  "net"
  "log"
  "google.golang.org/grpc"

  "github.com/dreae/erebus/lib/config"
  pb "github.com/dreae/erebus/protobuf"
)

func Init(conf *config.Config) {
  lis, err := net.Listen("tcp", ":27015")
  if err != nil {
    log.Fatalf("Failed to listen: %v", err)
  }

  s := grpc.NewServer()
  pb.RegisterMasterServer(s, NewServer(conf))
  go s.Serve(lis)
}
