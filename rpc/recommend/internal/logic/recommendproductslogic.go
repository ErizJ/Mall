package recommend

import (
	"context"
	"mall/recommend/internal/svc"
	"sync"
)

// Product模型简化版
type Product struct {
	ProductId int64
	Name      string
	Price     float32
	ImageUrl  string
}

// 假设有用户购买历史数据，格式：userID -> 购买过的商品ID列表
var userPurchaseHistory = map[string][]int64{
	"user1": {101, 102, 103},
	"user2": {102, 104},
	"user3": {101, 105},
}

// 假设有商品简单信息表
var productDB = map[int64]Product{
	101: {101, "手机", 2999.99, "url1"},
	102: {102, "耳机", 199.99, "url2"},
	103: {103, "充电宝", 99.99, "url3"},
	104: {104, "手机壳", 49.99, "url4"},
	105: {105, "数据线", 29.99, "url5"},
}

type RecommendLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
}

func NewRecommendLogic(ctx context.Context, svcCtx *svc.ServiceContext) *RecommendLogic {
	return &RecommendLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
	}
}

// 简单基于用户购买的协同过滤推荐
func (l *RecommendLogic) RecommendProducts(userID string, limit int) ([]Product, error) {
	if limit <= 0 {
		limit = 10
	}

	userProducts, ok := userPurchaseHistory[userID]
	if !ok {
		// 新用户，推荐热门商品（简单逻辑）
		return l.recommendPopular(limit), nil
	}

	// 1. 找到购买相似的用户
	similarUsers := l.findSimilarUsers(userID, userProducts)

	// 2. 收集相似用户购买过但当前用户没买过的商品
	recommendSet := make(map[int64]struct{})
	userProductSet := make(map[int64]struct{})
	for _, pid := range userProducts {
		userProductSet[pid] = struct{}{}
	}
	for _, su := range similarUsers {
		for _, pid := range userPurchaseHistory[su] {
			if _, bought := userProductSet[pid]; !bought {
				recommendSet[pid] = struct{}{}
			}
		}
	}

	// 3. 转成列表返回
	var recs []Product
	for pid := range recommendSet {
		if p, exist := productDB[pid]; exist {
			recs = append(recs, p)
		}
		if len(recs) >= limit {
			break
		}
	}

	// 4. 不够则补充热门商品
	if len(recs) < limit {
		popular := l.recommendPopular(limit - len(recs))
		recs = append(recs, popular...)
	}

	return recs, nil
}

// 找购买相似的用户，简单Jaccard相似度
func (l *RecommendLogic) findSimilarUsers(userID string, userProducts []int64) []string {
	threshold := 0.3
	var similar []string

	userSet := make(map[int64]struct{})
	for _, pid := range userProducts {
		userSet[pid] = struct{}{}
	}

	var mu sync.Mutex

	for otherUser, products := range userPurchaseHistory {
		if otherUser == userID {
			continue
		}
		intersectionCount := 0
		unionSet := make(map[int64]struct{})
		for pid := range userSet {
			unionSet[pid] = struct{}{}
		}
		for _, pid := range products {
			unionSet[pid] = struct{}{}
			if _, ok := userSet[pid]; ok {
				intersectionCount++
			}
		}

		jaccard := float64(intersectionCount) / float64(len(unionSet))
		if jaccard >= threshold {
			mu.Lock()
			similar = append(similar, otherUser)
			mu.Unlock()
		}
	}
	return similar
}

// 简单热门商品推荐
func (l *RecommendLogic) recommendPopular(limit int) []Product {
	// 这里直接返回productDB前几个，实际可用销量排行等
	var popular []Product
	count := 0
	for _, p := range productDB {
		popular = append(popular, p)
		count++
		if count >= limit {
			break
		}
	}
	return popular
}
