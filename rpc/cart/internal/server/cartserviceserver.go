// Code generated by goctl. DO NOT EDIT.
// goctl 1.8.4
// Source: cart.proto

package server

import (
	"context"

	"mall/cart/internal/logic"
	"mall/cart/internal/svc"
	"mall/cart/mall/cart"
)

type CartServiceServer struct {
	svcCtx *svc.ServiceContext
	cart.UnimplementedCartServiceServer
}

func NewCartServiceServer(svcCtx *svc.ServiceContext) *CartServiceServer {
	return &CartServiceServer{
		svcCtx: svcCtx,
	}
}

func (s *CartServiceServer) AddItem(ctx context.Context, in *cart.AddItemRequest) (*cart.Empty, error) {
	l := logic.NewAddItemLogic(ctx, s.svcCtx)
	return l.AddItem(in)
}

func (s *CartServiceServer) RemoveItem(ctx context.Context, in *cart.RemoveItemRequest) (*cart.Empty, error) {
	l := logic.NewRemoveItemLogic(ctx, s.svcCtx)
	return l.RemoveItem(in)
}

func (s *CartServiceServer) ListItems(ctx context.Context, in *cart.ListItemsRequest) (*cart.ListItemsResponse, error) {
	l := logic.NewListItemsLogic(ctx, s.svcCtx)
	return l.ListItems(in)
}

func (s *CartServiceServer) ClearCart(ctx context.Context, in *cart.ClearCartRequest) (*cart.Empty, error) {
	l := logic.NewClearCartLogic(ctx, s.svcCtx)
	return l.ClearCart(in)
}
