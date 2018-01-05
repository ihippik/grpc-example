package main

import (
	pb "github.com/ihippik/grpc-test/protocol"
	"golang.org/x/net/context"
	"net"
	"fmt"
	"google.golang.org/grpc"
)
type server struct{}

func (s *server) Get(ctx context.Context, in *pb.GetUserRequest)(*pb.User, error){
	fmt.Println(in.Id)
	switch in.Id{
	case 1:
		return &pb.User{Name:"Вася", Email:"loop@lop.lp"}, nil
		break
	default:
		return &pb.User{Name:"Петя", Email:"test@test.ts"}, nil
		break
	}
	return &pb.User{Name:"Вася", Email:"loop@lop.lp"}, nil
}

func main() {
	l, err:= net.Listen("tcp", ":55555")
	if err !=nil{
		fmt.Println(err.Error())
	}
	s:= grpc.NewServer()
	pb.RegisterGetUserServer(s,&server{})
	s.Serve(l)
}