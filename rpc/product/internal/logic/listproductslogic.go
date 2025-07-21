package logic

import (
	"context"

	"mall/rpc/product/internal/dao"
	"mall/rpc/product/internal/svc"
	product "mall/rpc/product/pb"

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

	// 1. 构造 DAO 请求参数
	req := &dao.ListProductsRequest{
		Page:        in.Page,
		PageSize:    in.PageSize,
		Keyword:     in.Keyword,
		CategoryIds: in.CategoryIds,
	}

	// 2. 调用 DAO 查询数据库
	daoResp, err := dao.ListProducts(l.svcCtx.DB, req)
	if err != nil {
		return nil, err
	}

	// 3. 类型转换为 RPC 返回格式
	resp := &product.ListProductsResponse{
		Total:    daoResp.Total,
		Products: make([]*product.Product, 0, len(daoResp.Products)),
	}

	for _, p := range daoResp.Products {
		resp.Products = append(resp.Products, &product.Product{
			ProductId:   p.ProductId,
			Name:        p.Name,
			CategoryId:  p.CategoryId,
			Price:       p.Price,
			Description: p.Description,
			ImageUrl:    p.ImageUrl,
			Stock:       p.Stock,
		})
	}

	return resp, nil
}
