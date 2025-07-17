package logic

import (
	"context"

	"mall/cart/internal/svc"
	"mall/cart/mall/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ClearCartLogic) ClearCart(in *cart.ClearCartRequest) (*cart.Empty, error) {
	// todo: add your logic here and delete this line

	return &cart.Empty{}, nil
}
