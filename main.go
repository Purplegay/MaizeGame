package main

import (
	"flag"
	"log"
	"net"
	myproto "rpctest/proto"
	"rpctest/rpcserver"

	"google.golang.org/grpc"
)

const (
	port = "18888"
)

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", "0.0.0.0:"+port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	myproto.RegisterSumServer(s, new(rpcserver.RpcPlayer))
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
