package logic

import (
	"context"

	"mall/rpc/product/internal/svc"
	product "mall/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type UpdateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewUpdateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *UpdateProductLogic {
	return &UpdateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 更新商品
func (l *UpdateProductLogic) UpdateProduct(in *product.UpdateProductRequest) (*product.Empty, error) {
	// todo: add your logic here and delete this line

	return &product.Empty{}, nil
}
