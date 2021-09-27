package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/tugzera/go-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect to gRPC: %e", err)
	}

	defer connection.Close()

	client := pb.NewUserServiceClient(connection)

	// AddUser(client)
	// AddUserVerbose(client)
	AddUsers(client)

}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Bruno",
		Email: "damascenobdm@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not AddUser: %e", err)
	}

	fmt.Println(res)
}

func AddUserVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Bruno",
		Email: "damascenobdm@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)
	if err != nil {
		log.Fatalf("Could not AddUser: %e", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("Could not receive stream message: %e", err)
		}
		fmt.Println("Status: ", stream.Status)
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "123",
			Name:  "Bruno",
			Email: "damascenobdm@gmail.com",
		},
		&pb.User{
			Id:    "124",
			Name:  "Jose",
			Email: "jose@gmail.com",
		},
		&pb.User{
			Id:    "125",
			Name:  "Leo",
			Email: "leo@gmail.com",
		},
		&pb.User{
			Id:    "126",
			Name:  "Alice",
			Email: "alice@gmail.com",
		},
		&pb.User{
			Id:    "127",
			Name:  "Iza",
			Email: "iza@gmail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())
	if err != nil {
		log.Fatalf("Error sending stream %e", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("Error receiving response stream %e", err)
	}

	fmt.Println(res)

}
