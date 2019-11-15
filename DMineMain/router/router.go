package router

import (
	"fmt"
	"github.com/dMineMain/core"
	"github.com/dMineMain/service"
	"github.com/gin-gonic/gin"
	"github.com/golang/protobuf/proto"
	"github.com/proto/types"
	"io/ioutil"
	"net/http"
	"strconv"
)


type RegisterParam struct{
	Address string  `form:"address"`
	Weight int `form:"weight"`
}

func InitRouter() *gin.Engine{
	r := gin.Default()
	r.POST("/register", Register)
	r.GET("/add_piece",AddPiece)
	r.GET("/gen_post",GenPoSt)
	r.POST("/seal_result",SealResult)
	r.POST("/post_result",PostResult)
	return r
}

//http+protobuf形式
func SealResult(ctx *gin.Context){
	var sealReq types.SealResultReq
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	proto.Unmarshal(data,&sealReq)
	fmt.Println("seal param",sealReq)
	ctx.ProtoBuf(http.StatusOK,&types.SealResultRes{Got:true})
}
//http+protobuf形式
func PostResult(ctx *gin.Context){
	var postReq types.PoStResultReq
	data, _ := ioutil.ReadAll(ctx.Request.Body)
	proto.Unmarshal(data,&postReq)
	fmt.Println("post param",postReq)
	ctx.ProtoBuf(http.StatusOK,&types.PoStResultRes{Got:true})
}
//注册服务
func Register(ctx *gin.Context){
	var reg RegisterParam
	if err := ctx.ShouldBind(&reg); err != nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1001, "Msg":"参数错误"})
		return
	}
	err := core.Register(reg.Address)
	if err!=nil {
		ctx.JSON(http.StatusOK, gin.H{"Code": 1002, "Msg":"注册错误:"+err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"Code": 0, "Msg":"注册成功"})
}
//AddPiece test
func AddPiece(ctx *gin.Context){

	//for i:=0;i<30;i++{
		service.AddPiece(&types.AddPieceReq{Ip:[]byte(strconv.Itoa(1))})
	//}
}
//GenPoSt test
func GenPoSt(ctx *gin.Context){
	service.GenPoSt(&types.GenPoStReq{Ip:[]byte("ip"),Seed:[]byte("seed")})
}


