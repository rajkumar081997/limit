package main


import (
	"context"
	"log"
	"net"
	"time"

	"github.com/boltdb/bolt"
	pb "github.com/m/v2/server"
	"google.golang.org/grpc"
)
type server struct{

}
 var world=[]byte("world")
func main(){
	
	lis,err:=net.Listen("tcp",":8001")
	if err!=nil{
		log.Fatal(err)
	}

	
	ser:=grpc.NewServer()
    pb.RegisterGetItemServer(ser,&pb.UnimplementedGetItemServer{})
    ser.Serve(lis)
	

}

func (s *server) Item(ctx context.Context,request *pb.Store)(*pb.Store,error){
	db,er:=bolt.Open("../my.bolt", 0600, &bolt.Options{Timeout: 1 * time.Second})
	if(er!=nil){
		log.Fatal(er)
	}

	defer db.Close()
	key:=[]byte("num")
	value:=[]byte(request.S)

	er=db.Update(func (tx *bolt.Tx) error {
		bucket,er:=tx.CreateBucketIfNotExists(world)
		er=bucket.Put(key,value)
		if er!=nil{
			log.Fatal(er)
		}
       return &pb.Store{S:"Data stored Sucessfully"},nil

	}

}