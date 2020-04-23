package main

import (
	pb "goStudy/goProject20/userInfo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"log"
	"net"
)

const (
	port = ":50051"
)

type server struct{}

func (s *server) GetUserInfo(ctx context.Context, req *pb.UserInfoRequest) (*pb.MessageResponse, error) {
	//初始一个响应对象
	mess := &pb.MessageResponse{
		Code: "0x000000",
		Msg:  "请求成功",
		Data: &pb.UserInfo{
			Userid:   "0001",
			Username: "admin",
			Userage:  20,
		},
	}

	if req.Userid != "001" {
		mess.Data = &pb.UserInfo{}
	}

	return mess, nil
}

func main() {
	//创建一个tcp服务，暴露出来的接口是50051
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	//注册服务接口
	pb.RegisterUserServer(s, &server{})
	s.Serve(lis)
}
