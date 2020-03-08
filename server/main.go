package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"strings"

	"google.golang.org/grpc"

	demopb "grpc-demo/proto/v1"
)

type demoService struct {
	retry int
}

func (d demoService) TryUnary(ctx context.Context, req *demopb.HelloRequest) (*demopb.HelloResponse, error) {
	return &demopb.HelloResponse{
		Result: fmt.Sprintf("Hello, %s!", req.GetName()),
	}, nil
}

func (d demoService) TryServerStreaming(req *demopb.HelloRequest, stream demopb.DemoServices_TryServerStreamingServer) error {
	for i := 0; i < d.retry; i++ {
		if err := stream.Send(&demopb.HelloResponse{
			Result: fmt.Sprintf("round %d: Hello, %s!", i+1, req.GetName()),
		}); err != nil {
			return err
		}
	}
	return nil
}

func (d demoService) TryClientStreaming(stream demopb.DemoServices_TryClientStreamingServer) error {
	var req *demopb.HelloRequest
	var err error
	var names []string
	for {
		req, err = stream.Recv()
		if err == io.EOF {
			// client has already sent all messages
			return stream.SendAndClose(&demopb.HelloResponse{
				Result: fmt.Sprintf("Hello, %s!", strings.Join(names, ", ")),
			})
		}
		names = append(names, req.GetName())
	}
}

func (d demoService) TryBidirectionalStreaming(stream demopb.DemoServices_TryBidirectionalStreamingServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		if err = stream.Send(&demopb.HelloResponse{
			Result: fmt.Sprintf("Hello, %s!", in.GetName()),
		}); err != nil {
			return err
		}
	}
}

var _ demopb.DemoServicesServer = demoService{}

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", 32303))
	if err != nil {
		panic(err)
	}

	server := grpc.NewServer()
	demopb.RegisterDemoServicesServer(server, demoService{retry: 3})
	log.Printf("%#v", server.GetServiceInfo())
	server.Serve(lis)
}
