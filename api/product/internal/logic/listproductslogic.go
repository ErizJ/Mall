package logic

import (
	"context"

	"mall/api/product/internal/svc"
	"mall/api/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewListProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListProductsLogic {
	return &ListProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *ListProductsLogic) ListProducts(req *types.ListProductsRequest) (resp *types.ListProductsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
