package logic

import (
	"context"

	"mall/cart/internal/svc"
	"mall/cart/mall/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type AddItemLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewAddItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *AddItemLogic {
	return &AddItemLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *AddItemLogic) AddItem(in *cart.AddItemRequest) (*cart.Empty, error) {
	// todo: add your logic here and delete this line

	return &cart.Empty{}, nil
}
