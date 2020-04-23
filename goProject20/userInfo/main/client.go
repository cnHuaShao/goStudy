package main

import (
	"log"
	"os"

	pb "goStudy/goProject20/userInfo"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

const (
	address       = "localhost:50051"
	defaultUserId = "001"
)

func main() {
	//建立与服务端的连接
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("did not connect: %v", err)
	}
	defer conn.Close()

	//准备好客户端
	c := pb.NewUserClient(conn)
	//使用默认值
	userId := defaultUserId
	//如果有输入值，则使用输入值，无则使用默认值
	if len(os.Args) > 1 {
		userId = os.Args[1]
	}
	//调用获取用户信息接口
	r, err := c.GetUserInfo(context.Background(), &pb.UserInfoRequest{
		Userid: userId,
	})
	if err != nil {
		log.Fatal("未获取到用户信息: %v", err)
	}
	log.Printf("mess: %s", r)
}
