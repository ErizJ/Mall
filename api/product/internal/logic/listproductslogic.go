package logic

import (
	"context"
	"fmt"

	"mall/api/product/internal/svc"
	"mall/api/product/internal/types"
	product "mall/rpc/product/pb"

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
	rpcReq := &product.ListProductsRequest{
		Page:        req.Page,
		PageSize:    req.PageSize,
		Keyword:     req.Keyword,
		CategoryIds: req.CategoryIds,
	}

	rpcResp, err := l.svcCtx.ProductClient.ListProducts(context.Background(), rpcReq)
	if err != nil {
		l.Logger.Error(fmt.Sprintf("Failed to get product list: %v", err))
		return nil, err
	}

	resp = &types.ListProductsResponse{
		Total:    rpcResp.Total,
		Products: make([]types.Product, 0, len(rpcResp.Products)),
	}

	for _, p := range rpcResp.Products {
		resp.Products = append(resp.Products, types.Product{
			ProductId:   p.ProductId,
			Name:        p.Name,
			CategoryId:  p.CategoryId,
			Price:       p.Price,
			Description: p.Description,
			ImageUrl:    p.ImageUrl,
			Stock:       p.Stock,
			IsAvailable: p.IsAvailable,
		})
	}

	return resp, nil
}
