package main

import (
	"flag"
	"fmt"
	"go-servers/user_server/handler"
	"go-servers/user_server/initialize"
	"go-servers/user_server/proto"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"net"
)

func main() {

	IP := flag.String("ip", "0.0.0.0", "ip地址")
	Port := flag.Int("port", 50051, "ip地址")
	// 初始化
	initialize.InitLogger()
	initialize.InitConfig()
	initialize.InitDB()

	flag.Parse()

	zap.S().Info("ip:", *IP)

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
