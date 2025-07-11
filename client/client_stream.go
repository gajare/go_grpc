package main

import (
	"context"
	pb "go_grpc/proto"
	"log"
	"time"
)

func callSayHelloClientStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Client straming started...")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		log.Fatalf("Cloud not send me :%v", err)
	}
	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		log.Printf("sent the request with name :%s", name)
		time.Sleep(2 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	log.Printf("Client streaming finished")
	if err != nil {
		log.Fatalf("Error while receiving :%v", err)
	}
	log.Printf("%v", res.Messages)
}
