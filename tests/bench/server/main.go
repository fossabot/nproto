package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/golang/protobuf/ptypes/empty"
	"github.com/huangjunwen/nproto/nproto/nprpc"
	"github.com/huangjunwen/nproto/nproto/taskrunner"
	"github.com/nats-io/go-nats"

	benchapi "github.com/huangjunwen/nproto/tests/bench/api"
)

type Bench struct{}

func (svc Bench) Nop(ctx context.Context, input *empty.Empty) (output *empty.Empty, err error) {
	return input, nil
}

func (svc Bench) Echo(ctx context.Context, input *benchapi.EchoMsg) (output *benchapi.EchoMsg, err error) {
	return input, nil
}

var (
	serverNum      int
	maxConcurrency int
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			log.Fatal(err)
		}
	}()

	flag.IntVar(&serverNum, "s", 10, "Server number")
	flag.IntVar(&maxConcurrency, "x", 5000, "Max concurrency")
	flag.Parse()

	runner := taskrunner.NewLimitedRunner(maxConcurrency, -1)
	defer runner.Close()

	log.Printf("Launching %d server.\n", serverNum)
	for i := 0; i < serverNum; i++ {
		nc, err := nats.Connect(
			nats.DefaultURL,
			nats.MaxReconnects(-1),
			nats.Name(fmt.Sprintf("server-%d-%d", os.Getpid(), i)),
		)
		if err != nil {
			panic(err)
		}
		defer nc.Close()

		server, err := nprpc.NewNatsRPCServer(
			nc,
			nprpc.ServerOptTaskRunner(runner),
		)
		if err != nil {
			panic(err)
		}
		defer server.Close()

		if err := benchapi.ServeBench(server, benchapi.SvcName, Bench{}); err != nil {
			panic(err)
		}
	}

	log.Printf("Launched.\n")

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c
}
