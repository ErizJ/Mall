package dao

import (
	"strings"

	"mall/rpc/product/internal/dao/model"

	"gorm.io/gorm"
)

type ListProductsRequest struct {
	Page        int64
	PageSize    int64
	Keyword     string
	CategoryIds []int64
}

type ListProductsResponse struct {
	Total    int64
	Products []*model.Product
}

// 列出商品
func ListProducts(db *gorm.DB, req *ListProductsRequest) (*ListProductsResponse, error) {

	var total int64
	var products []*model.Product

	query := db.Model(&model.Product{})

	if strings.TrimSpace(req.Keyword) != "" {
		query = query.Where("name LIKE ?", "%"+strings.TrimSpace(req.Keyword)+"%")
	}

	if len(req.CategoryIds) > 0 {
		query = query.Where("category_id IN ?", req.CategoryIds)
	}

	// 获取总数
	if err := query.Count(&total).Error; err != nil {
		return nil, err
	}

	offset := (req.Page - 1) * req.PageSize
	err := query.Order("product_id DESC").Offset(int(offset)).Limit(int(req.PageSize)).Find(&products).Error
	if err != nil {
		return nil, err
	}

	return &ListProductsResponse{
		Total:    total,
		Products: products,
	}, nil
}
