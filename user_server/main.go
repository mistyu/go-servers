package main

import (
	"flag"
	"fmt"
	"go-servers/user_server/handler"
	"go-servers/user_server/proto"
	"google.golang.org/grpc"
	"net"
)

func main() {

	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("prot", 50051, "ip地址")

	flag.Parse()

	server := grpc.NewServer()
	proto.RegisterUserServer(server, &handler.UserServer{})
	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", *IP, *Port))
	if err != nil {
		panic("failed to listen:" + err.Error())
	}
	err = server.Serve(lis)
	if err != nil {
		panic("failed to grpc:" + err.Error())
	}

}
