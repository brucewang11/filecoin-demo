package router

import (
	"github.com/dMineMain/service"
	"github.com/gin-gonic/gin"
	"github.com/proto/types"
	"strconv"
)


type RegisterParam struct{
	Address string  `form:"address"`
	Weight int `form:"weight"`
}

func InitRouter() *gin.Engine{
	r := gin.Default()
	r.GET("/seal_pre_commit",SealPreCommit)
	r.GET("/seal_commit",SealCommit)
	r.GET("/gen_candidates",GenCandidates)
	r.GET("/gen_post",GenPoSt)
	return r
}
//AddPiece test
func SealPreCommit(ctx *gin.Context){
	for i:=0;i<30;i++{
		service.SealPreCommit(&types.PreCommitReq{ProverID:[]byte(strconv.Itoa(i))})
	}
}

func SealCommit(ctx *gin.Context){
	for i:=0;i<30;i++{
		service.SealCommit(&types.SealCommitReq{ProverID:[]byte(strconv.Itoa(i))})
	}
}

func GenCandidates(ctx *gin.Context){
	service.GenCandidates(&types.CandidatesReq{ProverID:[]byte(strconv.Itoa(1))})
}

func GenPoSt(ctx *gin.Context){
	service.GenPoSt(&types.GenPoStReq{Ip:[]byte("ip"),Seed:[]byte("seed")})
}


