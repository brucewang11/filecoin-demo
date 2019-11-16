package service

import (
	"context"
	"errors"
	"github.com/dMineMain/core"
	"github.com/proto/register"
)

type RegisterServiceServer struct {}

//注册服务
func (s *RegisterServiceServer)Register(ctx context.Context, in *register.RegisterReq) (*register.RegisterRes, error) {
	if in.Ip ==""{
		return nil,errors.New("注册参数错误")
	}
	err := core.Register(in.Ip)
	if err!=nil {
		return nil,errors.New("注册错误"+err.Error())
	}
	return &register.RegisterRes{Code:0}, nil
}
