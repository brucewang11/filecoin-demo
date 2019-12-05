package service

import (
	"context"
	"github.com/proto/types"
)

type DMineSubServer struct {}


func (s *DMineSubServer) SealPreResult(ctx context.Context, in *types.PreResultReq) (*types.PreResultRes, error) {
	return &types.PreResultRes{Got:true}, nil
}

func (s *DMineSubServer) SealResult(ctx context.Context, in *types.SealResultReq) (*types.SealResultRes, error) {
	return &types.SealResultRes{Got:true}, nil
}

func (s *DMineSubServer) CanResult(ctx context.Context, in *types.CanResultReq) (*types.CanResultRes, error) {
	return &types.CanResultRes{Got:true}, nil
}

func (s *DMineSubServer) PoStResult(ctx context.Context, in *types.PoStResultReq) (*types.PoStResultRes, error) {
	return &types.PoStResultRes{Got:true}, nil
}




