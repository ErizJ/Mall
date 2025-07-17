package logic

import (
	"context"

	"mall/api/product/internal/svc"
	"mall/api/product/internal/types"

	"github.com/zeromicro/go-zero/core/logx"
)

type SearchSuggestionsLogic struct {
	logx.Logger
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewSearchSuggestionsLogic(ctx context.Context, svcCtx *svc.ServiceContext) *SearchSuggestionsLogic {
	return &SearchSuggestionsLogic{
		Logger: logx.WithContext(ctx),
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

func (l *SearchSuggestionsLogic) SearchSuggestions(req *types.SuggestionsRequest) (resp *types.SuggestionsResponse, err error) {
	// todo: add your logic here and delete this line

	return
}
