package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/shumbo/grpc-web-error-details/sample/proto"
	"google.golang.org/genproto/googleapis/rpc/errdetails"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

const (
	address     = "localhost:50051"
	defaultName = "world"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewSampleServiceClient(conn)

	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
	r, err = c.SayError(ctx, &pb.ErrorRequest{Code: 500})
	if err != nil {
		st := status.Convert(err)
		fmt.Printf("%v, %s", st, st)
		for _, detail := range st.Details() {
			switch errorType := detail.(type) {
			case *errdetails.BadRequest:
				violations := errorType.GetFieldViolations()
				for _, v := range violations {
					fmt.Printf("error: %s, detail: %s\n", v.Field, v.Description)
				}

			default:
				fmt.Printf("default: %s\n", errorType)

			}
		}
		log.Fatalf("could not say error: %v", err)
		return
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
