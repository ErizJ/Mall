package logic

import (
	"context"

	"mall/product/internal/svc"
	"mall/product/mall/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 商品列表
func (l *ListProductsLogic) ListProducts(in *product.ListProductsRequest) (*product.ListProductsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.ListProductsResponse{}, nil
}
