package logic

import (
	"context"

	"mall/api/order/internal/svc"
	"mall/api/order/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateOrderStatusLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewUpdateOrderStatusLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateOrderStatusLogic {
	return &UpdateOrderStatusLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *UpdateOrderStatusLogic) UpdateOrderStatus(req *types.UpdateOrderStatusRequest) (resp *types.Empty, err error) {
	// todo: add your logic here and delete this line

	return
}
