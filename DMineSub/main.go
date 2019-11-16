package main

import (
	"context"
	"fmt"
	"github.com/dMineSub/common"
	"github.com/dMineSub/service"
	"github.com/dMineSub/utils"
	"github.com/proto/register"
	pb "github.com/proto/services"
	"google.golang.org/grpc"
	"net"
	"os"
	"strings"
	"time"
)

var (
	err1 = "缺少参数， 1.本服务的grpc adadress，2.主节点的注册地址，如命令 ./dMineSub 127.0.0.1:10001 127.0.0.1:5001"
	err2 = "第一个参数格式错误"
	err3 = "第二个参数格式错误"
)
type resData struct {
	Code int
	Msg string
}

func main() {
	if len(os.Args)!=3{
		panic(err1)
	}
	grpcAddress := os.Args[1]
	registerAddress := os.Args[2]
	if !utils.VerifyIp(grpcAddress){
		panic(err2)
	}
	if !utils.VerifyIp(registerAddress){
		panic(err3)
	}
	//初始化url路径
	common.InitUrl(registerAddress)

	//初始化子节点连接主节点的grpc
	err := service.InitGrpcClient()
	if err!=nil {
		panic("子节点连接主节点grpc错误"+err.Error())
	}

	strPort := strings.Split(grpcAddress,":")[1]
	lis, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDMineMainServer(grpcServer, &service.DMineMainServer{})

	//启动子节点grpc
	go grpcServer.Serve(lis)

	time.Sleep(2*time.Second)
	//2秒后向主节点注册服务
	_,err=service.RegisterClient.Register(context.Background(),&register.RegisterReq{Ip:grpcAddress})
	if err != nil {
		panic(err)
	}
	fmt.Println("注册成功",grpcAddress)
	service.ReadChannel()
}



