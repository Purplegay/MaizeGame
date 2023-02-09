package rpcserver

import (
	"context"
	"log"
	myproto "rpctest/proto"
)

type RpcPlayer struct {
	myproto.UnimplementedSumServer
}

// SayHello implements helloworld.GreeterServer
func (this *RpcPlayer) SayHello(ctx context.Context, in *myproto.SumNumsREQ) (*myproto.SumNumsACK, error) {
	log.Printf("Received: %v", in.Nums)
	return &myproto.SumNumsACK{Result: in.Nums * 10}, nil
}
