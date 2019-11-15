package service

import (
	"bytes"
	"context"
	"fmt"
	"github.com/dMineSub/common"
	"github.com/golang/protobuf/proto"
	types "github.com/proto/types"
	"io/ioutil"
	"net/http"
)

type DMineMainServer struct {}

var sealChannel = make(chan types.AddPieceReq, 1000)
var postChannel = make(chan types.GenPoStReq, 1000)

//读取并处理channel数据
func ReadChannel (){
	var piece types.AddPieceReq
	var genPoSt types.GenPoStReq
	for{
		select {
		case piece = <-sealChannel:
			SealResult(piece)
		case genPoSt = <-postChannel :
			PoStResult(genPoSt)
		}
	}
}


func (s *DMineMainServer) AddPiece(ctx context.Context, in *types.AddPieceReq) (*types.AddPieceRes, error) {
	fmt.Println("addpiece")
	sealChannel <- *in //防止任务并发，写进channel
	return &types.AddPieceRes{SectorId: 1}, nil

}
func (s *DMineMainServer) GenPoSt(ctx context.Context, in *types.GenPoStReq) (*types.GenPoStRes, error) {
	fmt.Println("genpost")
	postChannel <- *in //防止任务并发，写进channel
	return &types.GenPoStRes{}, nil
}



func SealResult(in types.AddPieceReq){
	fmt.Println("SealResult任务完成,开始通知主节点",in)
	req := &types.SealResultReq{SectorID:11}
	protoReq,err := proto.Marshal(req)
	resp,err := http.Post(common.SealResulUrl,"Content-Type: application/x-protobuf",bytes.NewReader(protoReq))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("SealResult,http调用错误",err)
		// handle error
	}
	defer resp.Body.Close()
	res:= types.SealResultRes{}
	err = proto.Unmarshal(body,&res)
	if err!=nil {
		fmt.Println("SealResult,json格式错误",err)
	}
	fmt.Println("SealResult通知完成",res)
}

func PoStResult(in types.GenPoStReq){
	fmt.Println("PoStResult任务完成,开始通知主节点",in)
	req := &types.PoStResultReq{Proof:[]byte("proof")}
	protoReq,err := proto.Marshal(req)

	resp,err := http.Post(common.PostResultUrl,"Content-Type: application/x-protobuf",bytes.NewReader(protoReq))
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("PoStResult,http调用错误",err)
		// handle error
	}
	defer resp.Body.Close()
	res:= types.PoStResultRes{}
	err = proto.Unmarshal(body,&res)
	if err!=nil {
		fmt.Println("PoStResult,json格式错误",err)
	}
	fmt.Println("PoStResult通知完成",res)
}