package main

import (
	"context"
	"fmt"
	"log"
	"bufio"
	"os"

	pb "github.com/m/v2/server"
	"google.golang.org/grpc"
)

func main() {
	var st string
	input:=bufio.NewScanner(os.Stdin)
	for input.Scan(){
		st=input.Text()
	}
	con, err := grpc.Dial("localhost:8081")

	if err != nil {
		log.Fatal(err)
	}
	clt := pb.NewGetItemClient(con)
	resp, er := clt.Item(context.Background(), &pb.Store{S: st})
	if er != nil {
		log.Fatal(er)
	}

	fmt.Println(resp.S)

}
