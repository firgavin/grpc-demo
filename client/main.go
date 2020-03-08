package main

import (
	"context"
	"fmt"
	"google.golang.org/grpc"
	demopb "grpc-demo/proto/v1"
	"io"
	"log"
)

func main() {
	conn, err := grpc.Dial(fmt.Sprintf(":%d", 32303), grpc.WithInsecure())
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := demopb.NewDemoServicesClient(conn)

	{
		unaryResp, err := cli.TryUnary(context.Background(), &demopb.HelloRequest{
			Name: "name1",
		})
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(unaryResp.String())
	}

	{
		serverStreamingCli, err := cli.TryServerStreaming(context.Background(), &demopb.HelloRequest{
			Name: "name1",
		})
		if err != nil {
			log.Fatalln(err)
		}
		for {
			resp, err := serverStreamingCli.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalln(err)
			}
			log.Println("serverStreamRecv", resp)
		}
	}

	{
		stream, err := cli.TryClientStreaming(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		for i := 0; i < 10; i++ {
			_ = stream.Send(&demopb.HelloRequest{
				Name: fmt.Sprintf("name%d", i),
			})
		}
		resp, err := stream.CloseAndRecv()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println("clientStream closeAndRecv", resp)
	}

	{
		stream, err := cli.TryBidirectionalStreaming(context.Background())
		if err != nil {
			log.Fatalln(err)
		}
		done := make(chan struct{})
		go func() {
			for {
				in, err := stream.Recv()
				if err == io.EOF {
					// read done.
					close(done)
					return
				}
				if err != nil {
					log.Fatalln(err)
				}
				log.Println("bidirectional from server", in)
			}
		}()
		for i := 0; i < 10; i++ {
			_ = stream.Send(&demopb.HelloRequest{
				Name: fmt.Sprintf("name%d", i),
			})
		}
		stream.CloseSend()
		<-done
	}
}
