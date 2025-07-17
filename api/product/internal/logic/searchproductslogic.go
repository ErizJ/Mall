package logic

import (
	"context"

	"mall/api/product/internal/svc"
	"mall/api/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchProductsLogic {
	return &SearchProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchProductsLogic) SearchProducts(req *types.SearchRequest) (resp *types.SearchResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
