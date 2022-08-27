package main

import (
	"context"
	"fmt"
	"grpc-client-server/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalln("")
	}
	defer conn.Close()

	client := pb.NewFileServicesClient(conn)
	callListFiles(client)
}

func callListFiles(client pb.FileServicesClient) {
	res, err := client.ListFiles(context.Background(), &pb.ListFilesRequest{})
	if err != nil {
		log.Fatalln("aa")
	}
	fmt.Println(res.GetFilenames())

}
