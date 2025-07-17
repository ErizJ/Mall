package logic

import (
	"context"

	"mall/user/internal/svc"
	"mall/user/mall/user"

	"github.com/zeromicro/go-zero/core/logx"
)

type ListUsersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewListUsersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ListUsersLogic {
	return &ListUsersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 列出用户（管理员用）
func (l *ListUsersLogic) ListUsers(in *user.ListUsersRequest) (*user.ListUsersResponse, error) {
	// todo: add your logic here and delete this line

	return &user.ListUsersResponse{}, nil
}
