package logic

import (
	"context"

	"mall/api/api/cart/internal/svc"
	"mall/api/api/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RemoveItemLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRemoveItemLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RemoveItemLogic {
	return &RemoveItemLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RemoveItemLogic) RemoveItem(req *types.RemoveItemRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
