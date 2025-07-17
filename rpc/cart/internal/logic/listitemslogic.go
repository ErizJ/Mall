package logic

import (
	"context"

	"mall/cart/internal/svc"
	"mall/cart/mall/cart"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListItemsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListItemsLogic {
	return &ListItemsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

func (l *ListItemsLogic) ListItems(in *cart.ListItemsRequest) (*cart.ListItemsResponse, error) {
	// todo: add your logic here and delete this line

	return &cart.ListItemsResponse{}, nil
}
