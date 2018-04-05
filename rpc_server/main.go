package main

import (
	"flag"
	"log"
	"net"
	"context"

	"github.com/pasiukevich/rpc"
	pb "github.com/pasiukevich/rpc/proto"

	"fmt"
	"google.golang.org/grpc"
	"strconv"
)

func main() {
	portPtr := flag.String("port", "5000", "Port to listen.")
	backupPtr := flag.String("backup", "", "Path to file with backup in gob format. Used to restore previous state of server.")

	flag.Parse()

	port, err := strconv.Atoi(*portPtr)
	if err != nil {
		log.Fatalf("wrong port value: %v", *portPtr)
	}

	// create the data store
	dataStore := rpc.New()

	// try to restore data from file if it's given
	if *backupPtr != "" {
		dataStore.FromFile(*backupPtr)
	}

	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	fmt.Println("LUL", *portPtr)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterStorageServer(grpcServer, &rpcServer{})
	grpcServer.Serve(lis)

	log.Println("Server is listening on:", *portPtr)

}

type rpcServer struct{}

func (s *rpcServer) GetValue(ctx context.Context, req *pb.Request) (*pb.Response, error) {
	fmt.Println("LUL", req.Body)
	return &pb.Response{Body: "test"}, nil
}
