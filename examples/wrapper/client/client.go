package main

import (
	"context"

	"github.com/stack-labs/stack"
	"github.com/stack-labs/stack/client"
	proto "github.com/stack-labs/stack/examples/proto/service/rpc"
	log "github.com/stack-labs/stack/logger"
	"github.com/stack-labs/stack/pkg/metadata"
)

type logWrapper struct {
	client.Client
}

func (l *logWrapper) Call(ctx context.Context, req client.Request, rsp interface{}, opts ...client.CallOption) error {
	log.Infof("[Call] 请求服务：%s.%s", req.Service(), req.Endpoint())
	newMd, b := metadata.FromContext(ctx)
	if !b {
		newMd = metadata.Metadata{}
	}
	newMd["client-wrapped"] = "client-wrapped-value"
	ctx = metadata.NewContext(ctx, newMd)

	return l.Client.Call(ctx, req, rsp)
}

func NewClientWrapper() client.Wrapper {
	return func(cli client.Client) client.Client {
		return &logWrapper{cli}
	}
}

func main() {
	service := stack.NewService(
		stack.Name("wrap.client.cli"),
		stack.WrapClient(
			NewClientWrapper(),
		),
	)
	service.Init()

	cl := proto.NewGreeterService("wrap.client.service", service.Client())
	rsp, err := cl.Hello(context.Background(), &proto.HelloRequest{Name: "StackLabs"})
	if err != nil {
		panic(err)
	}

	log.Info(rsp.Greeting)
}
