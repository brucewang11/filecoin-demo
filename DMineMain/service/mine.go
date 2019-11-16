package service

import (
	"context"
	"fmt"
	"github.com/dMineMain/core"
	"github.com/proto/types"
)

func AddPiece(req *types.AddPieceReq){
	client,err := core.GetRandServer()
	if err!=nil {
		fmt.Println("addPiece错误",err)
		return
	}
	//todo 超时时间
	res,err := client.AddPiece(context.Background(),req)
	fmt.Println("addpiece返回值",res,err)

}

func GenPoSt(req *types.GenPoStReq){
	clients := core.GetAllServer()
	for _,v := range clients{
		//todo 超时时间
		v.GenPoSt(context.Background(),req)
	}
}
