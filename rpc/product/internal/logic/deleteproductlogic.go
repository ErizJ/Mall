package logic

import (
	"context"

	"mall/rpc/product/internal/svc"
	product "mall/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type DeleteProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewDeleteProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *DeleteProductLogic {
	return &DeleteProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 删除商品
func (l *DeleteProductLogic) DeleteProduct(in *product.DeleteProductRequest) (*product.Empty, error) {
	// todo: add your logic here and delete this line

	return &product.Empty{}, nil
}
