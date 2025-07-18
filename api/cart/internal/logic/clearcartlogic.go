package logic

import (
	"context"

	"mall/api/api/cart/internal/svc"
	"mall/api/api/cart/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ClearCartLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewClearCartLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ClearCartLogic {
	return &ClearCartLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ClearCartLogic) ClearCart(req *types.ClearCartRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
