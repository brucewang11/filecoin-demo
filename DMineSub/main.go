package main

import (
	"encoding/json"
	"fmt"
	"github.com/dMineSub/common"
	"github.com/dMineSub/service"
	"github.com/dMineSub/utils"
	pb "github.com/proto/services"
	"google.golang.org/grpc"
	"io/ioutil"
	"net"
	"net/http"
	"net/url"
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
	strPort := strings.Split(grpcAddress,":")[1]
	lis, err := net.Listen("tcp", ":"+strPort)
	if err != nil {
		fmt.Println("failed to listen: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterDMineMainServer(grpcServer, &service.DMineMainServer{})

	//启动grpc
	go grpcServer.Serve(lis)

	time.Sleep(2*time.Second)
	//2秒后向主节点注册

	resp, err := http.PostForm(common.RegisterUrl, url.Values{"address": {grpcAddress}, "weight": {"1"}})
	if err != nil {
		fmt.Println(err)
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
		// handle error
	}
	resp.Body.Close()
	res:= resData{}
	err = json.Unmarshal(body,&res)
	if err!=nil {
		fmt.Println("json格式错误",err)
		panic(err)
	}
	if res.Code!=0{
		panic(res.Msg)
	} else {
		fmt.Println(res.Msg)
	}
	service.ReadChannel()
}



