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

	r.GET("/add_piece",AddPiece)
	r.GET("/gen_post",GenPoSt)
	return r
}
//AddPiece test
func AddPiece(ctx *gin.Context){

	for i:=0;i<30;i++{
		service.AddPiece(&types.AddPieceReq{Ip:[]byte(strconv.Itoa(i))})
	}
}
//GenPoSt test
func GenPoSt(ctx *gin.Context){
	service.GenPoSt(&types.GenPoStReq{Ip:[]byte("ip"),Seed:[]byte("seed")})
}


