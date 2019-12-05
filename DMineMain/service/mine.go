package service

import (
	"context"
	"fmt"
	"github.com/dMineMain/core"
	"github.com/proto/types"
)

func SealPreCommit(req *types.PreCommitReq){
	client,err := core.GetRandServer()
	if err!=nil {
		fmt.Println("SealPreCommit",err)
		return
	}
	res,err := client.SealPreCommit(context.Background(),req)
	fmt.Println("SealPreCommit返回值",res,err)
}

func SealCommit(req *types.SealCommitReq){
	client,err := core.GetRandServer()
	if err!=nil {
		fmt.Println("SealCommit",err)
		return
	}
	res,err := client.SealCommit(context.Background(),req)
	fmt.Println("SealCommit 返回值",res,err)
}

func GenCandidates(req *types.CandidatesReq){
	clients := core.GetAllServer()
	for _,v := range clients{
		v.GenCandidates(context.Background(),req)
	}
}

func GenPoSt(req *types.GenPoStReq){
	clients := core.GetAllServer()
	for _,v := range clients{
		v.GenPoSt(context.Background(),req)
	}
}
