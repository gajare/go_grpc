package main

import (
	"context"
	pb "go_grpc/proto"
	"io"
	"log"
	"time"
)

func callHelloBidirectionalStream(client pb.GreetServiceClient, names *pb.NamesList) {
	log.Printf("Bi-Directional Streaming Started...")
	stream, err := client.SayHelloBidirectionalStreaming(context.Background())
	if err != nil {
		log.Fatalf("Could not send name : %v", err)
	}
	waitch := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				break
			}
			if err != nil {
				log.Fatalf("Error while straming: %v", err)
			}
			log.Println(message)
		}
		close(waitch)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			log.Fatalf("Error while sending: %v", err)
		}
		time.Sleep(2 * time.Second)
	}
	stream.CloseSend()
	<-waitch
	log.Printf("Bi-Direction streaming finished.")
}
