package svc

import (
	"mall/api/product/internal/config"
	"mall/rpc/product/pb"

	"github.com/zeromicro/go-zero/zrpc"
)

type ServiceContext struct {
	Config        config.Config
	ProductClient pb.ProductServiceClient
}

func NewServiceContext(c config.Config) *ServiceContext {
	conn := zrpc.MustNewClient(c.ProductRpc).Conn() // 获取 grpc.ClientConnInterface
	return &ServiceContext{
		Config:        c,
		ProductClient: pb.NewProductServiceClient(conn), // 正确调用构造函数
	}
}
