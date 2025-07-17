package logic

import (
	"context"

	"mall/api/recommend/internal/svc"
	"mall/api/recommend/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type RecommendProductsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendProductsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendProductsLogic {
	return &RecommendProductsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *RecommendProductsLogic) RecommendProducts(req *types.RecommendRequest) (resp *types.RecommendResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
