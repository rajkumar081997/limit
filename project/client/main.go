package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	pb "github.com/m/v2/server"
	"google.golang.org/grpc"
)

func main() {

	action := flag.String("action", " ", "no_one")
	input := flag.String("input", " ", "none")
	flag.Parse()

	con, err := grpc.Dial("localhost:8001", grpc.WithInsecure())

	if err != nil {
		log.Fatal(err)
	}
	clt := pb.NewGetItemClient(con)
	switch *action {
	case "store":
		rep, er := clt.Item(context.Background(), &pb.Store{S: *input})
		if er != nil {
			log.Fatal(er)
		}
		fmt.Println(rep.S)
	case "getid":
		rept, er := clt.GetId(context.Background(), &pb.Id{K: *input})
		if er != nil {
			log.Fatal(er)
		}
		fmt.Println(rept.S)
	case "list":
		rep, er := clt.List(context.Background(), &pb.Id{K: *input})
		if er != nil {
			log.Fatal(er)
		}
		fmt.Println(rep.S1)
	case "rm":
		rep, er := clt.Remove(context.Background(), &pb.Id{K: *input})
		if er != nil {
			log.Fatal(er)
		}
		fmt.Println(rep.S)
	}
}
