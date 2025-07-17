package logic

import (
	"context"

	"mall/product/internal/svc"
	"mall/product/mall/product"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSuggestionsLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewSearchSuggestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSuggestionsLogic {
	return &SearchSuggestionsLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 搜索提示词（联想词）
func (l *SearchSuggestionsLogic) SearchSuggestions(in *product.SuggestionsRequest) (*product.SuggestionsResponse, error) {
	// todo: add your logic here and delete this line

	return &product.SuggestionsResponse{}, nil
}
