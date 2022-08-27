package main

import (
	"context"
	"fmt"
	"grpc-client-server/pb"
	"io/ioutil"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedFileServicesServer
}

func (*server) ListFiles(ctx context.Context, req *pb.ListFilesRequest) (*pb.ListFilesRespnse, error) {
	fmt.Println("List files was invoked")

	dir := "/Users/coco/src/grpc-client-server/storage"
	pathes, err := ioutil.ReadDir(dir)
	if err != nil {
		return nil, err
	}

	var filenames []string
	for _, path := range pathes {
		if !path.IsDir() {
			filenames = append(filenames, path.Name())
		}
	}

	res := &pb.ListFilesRespnse{
		Filenames: filenames,
	}

	return res, nil

}

func main() {
	lis, err := net.Listen("tcp", "localhost:50051")
	if err != nil {
		log.Fatalln("err")
	}

	s := grpc.NewServer()
	pb.RegisterFileServicesServer(s, &server{})

	fmt.Println("server is running..")

	if err := s.Serve(lis); err != nil {
		log.Fatalln("hogehoge")
	}
}
