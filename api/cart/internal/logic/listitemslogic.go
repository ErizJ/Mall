package logic

import (
	"context"

	"mall/api/api/cart/internal/svc"
	"mall/api/api/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListItemsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListItemsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListItemsLogic {
	return &ListItemsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListItemsLogic) ListItems(req *types.ListItemsRequest) (resp *types.ListItemsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
