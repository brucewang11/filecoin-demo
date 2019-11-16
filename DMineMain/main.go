package main

import (
	"fmt"
	"github.com/dMineMain/router"
	"github.com/dMineMain/utils"
	"github.com/dMineMain/service"
	servicespb "github.com/proto/services"
	registerpb "github.com/proto/register"
	"google.golang.org/grpc"
	"net"
	"os"
)
var (
	err1 = "参数数量错误"
	err2 = "第一个参数格式错误,如： 5001"
	err3 = "第二个参数格式错误，如：5002"
)

func main() {
	if len(os.Args)==1{
		panic(err1)
	}
	port := os.Args[1]
	if !utils.VerifyPort(port){
		panic(err2)
	}
	//启动测试
	if len(os.Args) == 3 {
		r := router.InitRouter()
		go r.Run(fmt.Sprintf(":%s", os.Args[2]))
	}
	lis, err := net.Listen("tcp", ":"+port)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	//sealresult和postresult服务接口
	servicespb.RegisterDMineSubServer(grpcServer, &service.DMineSubServer{})
	//注册服务接口
	registerpb.RegisterRegisterServiceServer(grpcServer,&service.RegisterServiceServer{})
	//启动grpc
	fmt.Println("启动grpc")
	grpcServer.Serve(lis)

}
