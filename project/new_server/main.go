package main

import (
	"context"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"

	"github.com/boltdb/bolt"
	pb "github.com/m/v2/server"
	"google.golang.org/grpc"
)

var path = "../my.bolt"
var mp = map[string]string{}

type server struct{}

var world = []byte("world")

func main() {

	lis, err := net.Listen("tcp", ":8001")
	if err != nil {
		log.Fatal(err)
	}

	ser := grpc.NewServer()
	pb.RegisterGetItemServer(ser, &server{})
	ser.Serve(lis)

}

func (s *server) Item(ctx context.Context, request *pb.Store) (*pb.Store, error) {
	db, er := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if er != nil {
		log.Fatal(er)
	}

	defer db.Close()
	mp[request.Data] = strconv.Itoa(rand.Int())
	key := []byte(mp[request.Data])
	value := []byte(request.Data)

	er = db.Update(func(tx *bolt.Tx) error {
		bucket, er := tx.CreateBucketIfNotExists(world)
		if er != nil {
			log.Fatal(er)
		}
		er = bucket.Put(key, value)
		if er != nil {
			log.Fatal(er)
		}
		return er

	})
	sp := "id of data is " + mp[request.Data] + " Data Sucessfully Stored"
	return &pb.Store{Data: sp}, nil
}

func (s *server) GetId(clt context.Context, req *pb.Id) (*pb.Store, error) {
	db, er := bolt.Open(path, 0660, &bolt.Options{Timeout: 1 * time.Second})
	if er != nil {
		log.Fatal(er)
	}

	defer db.Close()
	var value string
	er = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(world)

		value = string(b.Get([]byte(req.Pick)))
		if er != nil {
			log.Println(er)
		}
		return er
	})

	return &pb.Store{Data: value}, nil
}

func (s *server) List(clt context.Context, req *pb.Id) (*pb.Group, error) {
	db, er := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()
	count, _ := strconv.Atoi(req.Pick)
	var s2 []string
	er = db.View(func(tx *bolt.Tx) error {
		b := tx.Bucket(world)

		cor := b.Cursor()
		var i = 0

		for k, v := cor.First(); k != nil; k, v = cor.Next() {
			s2 = append(s2, string(v))
			i++
			if i == count {
				break
			}
		}
		if er != nil {
			log.Println(er)
		}
		return er

	})

	return &pb.Group{Lst: s2}, nil

}

func (s *server) Remove(clt context.Context, req *pb.Id) (*pb.Store, error) {
	db, er := bolt.Open(path, 0600, &bolt.Options{Timeout: 1 * time.Second})
	if er != nil {
		log.Fatal(er)
	}
	defer db.Close()
	id := req.Pick

	er = db.Update(func(tx *bolt.Tx) error {
		b := tx.Bucket(world)
		b.Delete([]byte(id))
		if er != nil {
			log.Println(er)
		}
		return er
	})
	return &pb.Store{Data: "Data Deleted Sucessfully"}, nil
}
