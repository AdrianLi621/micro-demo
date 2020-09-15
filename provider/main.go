package main

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	_ "github.com/micro/go-plugins/registry/consul/v2"
	"provider/proto"
)

type Greeter struct {
}

func (g *Greeter) HelloTest(ctx context.Context, req *proto.SayRequest, resp *proto.SayResponse) error {
	fmt.Println("recived data")
	resp.Greeting = req.Name + "===="
	return nil
}

func main() {

	service := micro.NewService(micro.Name("greeter"))
	service.Init()

	proto.RegisterGreeterHandler(service.Server(), new(Greeter))

	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}
