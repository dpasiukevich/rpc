package main

import (
	"google.golang.org/grpc"

	"flag"
	"fmt"
	"context"

	pb "github.com/dpasiukevich/rpc/proto"
	"bufio"
	"os"
	"strings"
	"io"
	"log"
)


func main() {

	serverAddr := flag.String("serverAddr", "127.0.0.1:3000", "Address of rpc server.")

	flag.Parse()

	conn, err := grpc.Dial(*serverAddr, grpc.WithInsecure())
	if err != nil {
		fmt.Println("err creating rpc connection: ", err)
	}
	defer conn.Close()

	client := pb.NewStorageClient(conn)

	userInput := bufio.NewReader(os.Stdin)

	for {

		fmt.Print(*serverAddr, ">")
		inputString, err := userInput.ReadString('\n')
		if err != nil {
			fmt.Println(err)
		}

		tokens := strings.Split(inputString, " ")

		if len(tokens) == 0 {
			continue
		}

		command := tokens[0]

		ctx := context.Background()

		switch strings.TrimSpace(command) {
		case "get":
			resp, err := client.GetValue(ctx, &pb.GetValueRequest{Key: tokens[1]})
			if err != nil {
				fmt.Println(err)
				continue
			}
			fmt.Println(resp.Value)
		case "set":
			var ttl string
			if len(tokens) == 4 {
				ttl = tokens[3]
			} else {
				ttl = ""
			}
			_, err := client.SetValue(ctx, &pb.SetValueRequest{Key: tokens[1], Value: tokens[2], Ttl: ttl})
			if err != nil {
				fmt.Println(err)
				continue
			}
		case "get_keys":
			stream, err := client.GetKeys(ctx, &pb.GetKeysRequest{})
			if err != nil {
				fmt.Println(err)
				continue
			}
			for {
				key, err := stream.Recv()
				if err == io.EOF {
					break
				}
				if err != nil {
					log.Fatalf("%v.get_keys(_) = _, %v", client, err)
				}
				fmt.Println(key)
			}
		default:
			fmt.Println("unsupported command", command)
		}
	}
}