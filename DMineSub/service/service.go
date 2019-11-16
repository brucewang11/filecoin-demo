package service

import (

	"context"
	"fmt"
	"github.com/dMineSub/common"
	types "github.com/proto/types"
	"google.golang.org/grpc"
	registerpb "github.com/proto/register"
	servicespb "github.com/proto/services"
)

type DMineMainServer struct {}

var sealChannel = make(chan types.AddPieceReq, 1000)
var postChannel = make(chan types.GenPoStReq, 1000)


var RegisterClient registerpb.RegisterServiceClient
var dMineSubClient servicespb.DMineSubClient

func InitGrpcClient() error{
	conn, err := grpc.Dial(common.MainGrpcServerAddress, grpc.WithInsecure())
	if err!=nil {
		fmt.Println("子节点连接主节点grpc连接失败",err)
		return err
	}
	RegisterClient = registerpb.NewRegisterServiceClient(conn)
	dMineSubClient = servicespb.NewDMineSubClient(conn)
	return nil
}



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
	fmt.Println("addpiece",string(in.Ip))
	sealChannel <- *in //防止任务并发，写进channel
	return &types.AddPieceRes{SectorId: 1}, nil

}
func (s *DMineMainServer) GenPoSt(ctx context.Context, in *types.GenPoStReq) (*types.GenPoStRes, error) {
	fmt.Println("genpost")
	postChannel <- *in //防止任务并发，写进channel
	return &types.GenPoStRes{}, nil
}



func SealResult(in types.AddPieceReq){
	//todo 算法任务

	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.SealResult(context.Background(),&types.SealResultReq{SectorID:1})
	if err != nil {
		fmt.Println("SealResult,调用错误",err)
		// handle error
	}
	fmt.Println("SealResult通知完成",res)
}

func PoStResult(in types.GenPoStReq){
	//todo 算法任务


	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.PoStResult(context.Background(),&types.PoStResultReq{Proof:[]byte("proof")})
	if err != nil {
		fmt.Println("PoStResult,调用错误",err)
		// handle error
	}
	fmt.Println("SealResult通知完成",res)


	fmt.Println("PoStResult通知完成",res)
}