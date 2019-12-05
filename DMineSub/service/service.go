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

var preCommitChannel = make(chan types.PreCommitReq,1000)
var sealCommitChannel = make(chan types.SealCommitReq, 1000)
var candidatesChannel = make(chan types.CandidatesReq, 1000)
var genPoStChannel = make(chan types.GenPoStReq, 1000)


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
	var preCommit types.PreCommitReq
	var sealCommit types.SealCommitReq
	var candidates types.CandidatesReq
	var genPoSt types.GenPoStReq
	for{
		select {
		case preCommit = <-preCommitChannel:
			SealPreResult(preCommit)
		case sealCommit = <-sealCommitChannel :
			SealResult(sealCommit)
		case candidates = <-candidatesChannel :
			CanResult(candidates)
		case genPoSt = <-genPoStChannel :
			PoStResult(genPoSt)
		}
	}
}


func (s *DMineMainServer) SealPreCommit(ctx context.Context, in *types.PreCommitReq) (*types.PreCommitRes, error) {
	fmt.Println("SealPreCommit",string(in.ProverID))
	preCommitChannel <- *in //防止任务并发，写进channel
	return &types.PreCommitRes{}, nil

}

func (s *DMineMainServer) SealCommit(ctx context.Context, in *types.SealCommitReq) (*types.SealCommitRes, error) {
	fmt.Println("SealCommit",string(in.ProverID))
	sealCommitChannel <- *in //防止任务并发，写进channel
	return &types.SealCommitRes{}, nil
}

func (s *DMineMainServer) GenCandidates(ctx context.Context, in *types.CandidatesReq) (*types.CandidatesRes, error) {
	fmt.Println("GenCandidates",string(in.ProverID))
	candidatesChannel <- *in //防止任务并发，写进channel
	return &types.CandidatesRes{}, nil
}

func (s *DMineMainServer) GenPoSt(ctx context.Context, in *types.GenPoStReq) (*types.GenPoStRes, error) {
	fmt.Println("GenPoSt",string(in.Ip))
	genPoStChannel <- *in //防止任务并发，写进channel
	return &types.GenPoStRes{}, nil
}





func SealPreResult(in types.PreCommitReq){
	//todo 算法任务

	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.SealPreResult(context.Background(),&types.PreResultReq{})
	if err != nil {
		fmt.Println("SealPreResult,调用错误",err)
		// handle error
	}
	fmt.Println("SealPreResult 通知完成",res)
}


func SealResult(in types.SealCommitReq){
	//todo 算法任务

	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.SealResult(context.Background(),&types.SealResultReq{})
	if err != nil {
		fmt.Println("SealResult,调用错误",err)
		// handle error
	}
	fmt.Println("SealResult 通知完成",res)
}


func CanResult(in types.CandidatesReq){
	//todo 算法任务

	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.CanResult(context.Background(),&types.CanResultReq{})
	if err != nil {
		fmt.Println("CanResult,调用错误",err)
		// handle error
	}
	fmt.Println("CanResult 通知完成",res)
}


func PoStResult(in types.GenPoStReq){
	//todo 算法任务

	fmt.Println("算法任务结束，开始通知主节点")
	res,err := dMineSubClient.PoStResult(context.Background(),&types.PoStResultReq{Proof:[]byte("proof")})
	if err != nil {
		fmt.Println("PoStResult,调用错误",err)
		// handle error
	}
	fmt.Println("PoStResult 通知完成",res)

}