package main

import (
	"context"
	"log"
	"os"
	"os/signal"

	mathapi "github.com/huangjunwen/nproto/tests/math/api"

	"github.com/huangjunwen/nproto/nproto"
	"github.com/huangjunwen/nproto/nproto/nprpc"
	"github.com/nats-io/go-nats"
)

const (
	seqKey = "seq"
)

type Math struct{}

func (m Math) Sum(ctx context.Context, input *mathapi.SumRequest) (output *mathapi.SumReply, err error) {
	log.Printf("Got RPC call Sum(%v), seq: %s\n", input.Args, nproto.CurrRPCMetaData(ctx).Get(seqKey))
	reply := &mathapi.SumReply{}
	for _, arg := range input.Args {
		reply.Sum += arg
	}
	return reply, nil
}

func main() {
	nc, err := nats.Connect(nats.DefaultURL, nats.MaxReconnects(-1))
	if err != nil {
		log.Panic(err)
	}
	log.Printf("NATS connected.\n")
	defer nc.Close()

	server, err := nprpc.NewNatsRPCServer(nc)
	if err != nil {
		log.Panic(err)
	}
	log.Printf("NatsRPCServer created.\n")
	defer server.Close()

	if err := mathapi.ServeMath(server, mathapi.SvcName, Math{}); err != nil {
		log.Panic(err)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
