package logic

import (
	"context"

	"mall/rpc/product/internal/svc"
	product "mall/rpc/product/pb"

	"github.com/zeromicro/go-zero/core/logx"
)

type CreateProductLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewCreateProductLogic(ctx context.Context, svcCtx *svc.ServiceContext) *CreateProductLogic {
	return &CreateProductLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 创建商品
func (l *CreateProductLogic) CreateProduct(in *product.CreateProductRequest) (*product.Product, error) {
	// todo: add your logic here and delete this line

	return &product.Product{}, nil
}
