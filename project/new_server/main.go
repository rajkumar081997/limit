package main

import (
	"context"
	"log"
	"net"
	"time"

	store "github.com/m/v2/new_server/store"
	pb "github.com/m/v2/server"
	"go.uber.org/fx"
	"google.golang.org/grpc"
)

type server struct {
	st store.Store
}
func new_configure() store.Store {

	st, err := store.Newstore("bolt")
	if err != nil {
		log.Fatal(err)
	}
	return st
}
func Register(st store.Store, lc fx.Lifecycle) {
	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}
	serv := grpc.NewServer()
	pb.RegisterGetItemServer(serv, &server{st: st})
	serv.Serve(lis)
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			return nil
		},
		OnStop: func(context.Context) error {
			st.Close()
			return nil
		},
	})
}
func main() {
	app := fx.New(
		fx.Provide(new_configure),
		fx.Invoke(Register),
	)
	stopctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	if err := app.Stop(stopctx); err != nil {
		log.Fatal(err)
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
