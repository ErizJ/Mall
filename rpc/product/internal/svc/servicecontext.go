package svc

import (
	"mall/rpc/product/common"
	"mall/rpc/product/internal/config"

	"gorm.io/gorm"
)

type ServiceContext struct {
	Config config.Config
	DB     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
		DB:     common.InitDB(c.Mysql.Dsn), // 从配置中读取 DSN
	}
}
