package di

import (
	"github.com/google/wire"
	"github.com/Eric-WangHaitao/Go-0712/Week04/configs"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/dao"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/service"
	"github.com/Eric-WangHaitao/Go-0712/Week04/pkg/email"
)

// +build wireinject

// NewService 定义injector的函数签名
func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	wire.Build(service.NewService, email.EMailSet, dao.UserSet, configs.NewDb)
	return &service.Service{}, nil
}
