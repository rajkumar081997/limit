package main

import (
	"context"
	"log"
	"net"
	"os"
	"os/signal"
	"syscall"

	store "github.com/m/v2/new_server/store"
	pb "github.com/m/v2/server"
	"google.golang.org/grpc"
)

type server struct {
	st store.Store
}

func main() {
	st, err := store.Newstore("bolt")
	if err != nil {
		log.Fatal(err)
	}
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	ser := grpc.NewServer()
	pb.RegisterGetItemServer(ser, &server{st: st})
	ser.Serve(lis)

	sign:=make(chan os.Signal,2)
	signal.Notify(sign,os.Interrupt,syscall.SIGTERM)
	sgn:=<-sign
	if sgn==syscall.SIGTERM{
        st.Close()
	}

}

func (s *server) Item(ctx context.Context, req *pb.Store) (*pb.Store, error) {
	output, err := s.st.Item(req.Data)

	return &pb.Store{Data: output}, err

}

func (s *server) GetId(ctx context.Context, req *pb.Id) (*pb.Store, error) {
	output, err := s.st.GetId(req.Pick)
	return &pb.Store{Data: output}, err
}

func (s *server) List(ctx context.Context, req *pb.Id) (*pb.Group, error) {
	var output []string
	var err error
	output, err = s.st.List(req.Pick)
	return &pb.Group{Lst: output}, err
}

func (s *server) Remove(ctx context.Context, req *pb.Id) (*pb.Store, error) {
	output, err := s.st.Remove(req.Pick)

	return &pb.Store{Data: output}, err
}
