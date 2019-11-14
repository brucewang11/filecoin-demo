package service

import (
	"context"
	"fmt"
	types2 "github.com/proto/types"
)

type DMineMainServer struct {}

func (s *DMineMainServer) AddPiece(ctx context.Context, in *types2.AddPieceReq) (*types2.AddPieceRes, error) {
	fmt.Println("addpiece")
	return &types2.AddPieceRes{SectorId: 1}, nil

}
func (s *DMineMainServer) GenPoSt(ctx context.Context, in *types2.GenPoStReq) (*types2.GenPoStRes, error) {
	fmt.Println("genpost")
	return &types2.GenPoStRes{}, nil
}