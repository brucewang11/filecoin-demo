package main

import (
	"fmt"
	"github.com/dMineMain/router"
	"github.com/dMineMain/utils"
	"os"
)
var (
	err1 = "缺少参数"
	err2 = "第一个参数格式错误,如： 5001"
)

func main() {
	if len(os.Args)!=2{
		panic(err1)
	}
	port := os.Args[1]
	if !utils.VerifyPort(port){
		panic(err2)
	}
	r := router.InitRouter()
	r.Run(fmt.Sprintf(":%s", port))
}
