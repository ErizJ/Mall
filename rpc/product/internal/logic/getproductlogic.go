package logic

import (
	"context"

	"mall/rpc/product/internal/svc"
	product "mall/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type GetProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewGetProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *GetProductLogic {
	return &GetProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 获取商品详情
func (l *GetProductLogic) GetProduct(in *product.GetProductRequest) (*product.Product, error) {
	// todo: add your logic here and delete this line

	return &product.Product{}, nil
}
