package main

import (
	"flag"
	"log"
	"net"
	"context"

	"github.com/dpasiukevich/rpc"
	pb "github.com/dpasiukevich/rpc/proto"

	"fmt"
	"google.golang.org/grpc"
	"strconv"
	"strings"
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
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	rs := rpcServer{
		client: rpc.NewClient(dataStore),
	}

	grpcServer := grpc.NewServer()
	pb.RegisterStorageServer(grpcServer, &rs)

	log.Println("Server is listening on port:", port)

	grpcServer.Serve(lis)
}

type rpcServer struct{
	client *rpc.Client
}

func (s *rpcServer) GetValue(ctx context.Context, req *pb.GetValueRequest) (*pb.GetValueResponse, error) {
	reply, err := s.client.Exec("get", []string{strings.TrimSpace(req.Key)})
	if err != nil {
		return nil, err
	}
	return &pb.GetValueResponse{Value: reply}, nil
}

func (s *rpcServer) SetValue(ctx context.Context, req *pb.SetValueRequest) (*pb.SetValueResponse, error) {

	args := []string{
		strings.TrimSpace(req.Key),
		strings.TrimSpace(req.Value),
	}

	if req.Ttl != "" {
		args = append(args, strings.TrimSpace(req.Ttl))
	}

	_, err := s.client.Exec("set", args)
	if err != nil {
		return nil, err
	}
	return &pb.SetValueResponse{}, nil
}

func (s *rpcServer) GetKeys(req *pb.GetKeysRequest, stream pb.Storage_GetKeysServer) error {
	reply, err := s.client.Exec("keys", []string{})
	if err != nil {
		return err
	}

	for _, key := range strings.Split(reply, " ") {
		if err := stream.Send(&pb.GetKeysResponse{Key: key}); err != nil {
			return err
		}
	}
	return nil
}
