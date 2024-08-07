package main

import (
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "hello-grpc-go/proto"
)

func main() {

	// 尝试获取问候消息
	message, err := getGreeting("world")
	if err != nil {
		log.Fatalf("Error while greeting: %v", err)
	}
	log.Printf("Greeting: %s", message)
}

// getGreeting 连接到GRPC服务并发送Hello请求
func getGreeting(name string) (string, error) {

	client, err := grpc.NewClient("localhost:9090", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	conn := pb.NewHelloServiceClient(client)

	return Greeting(conn, name)
}

func Greeting(client pb.HelloServiceClient, name string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()

	response, err := client.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		return "", err
	}

	return response.Message, nil
}
