package main

import (
	"google.golang.org/grpc"

	"flag"
	"fmt"
	"context"

	pb "github.com/pasiukevich/rpc/proto"
)


func main() {

	serverAddr := flag.String("serverAddr", "127.0.0.1:5000", "Address of rpc server.")

	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err creating rpc connection: ", err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)

	resp, err := client.GetValue(context.Background(), &pb.Request{Body: "123"})
	if err != nil {
		fmt.Println("err getting response: ", err)
	}

	fmt.Println("response", resp.Body)
}
