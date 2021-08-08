//go:generate wire
//+build !wireinject

package di

import (
	"github.com/Eric-WangHaitao/Go-0712/Week04/configs"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/dao"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/service"
	"github.com/Eric-WangHaitao/Go-0712/Week04/pkg/email"
)

func NewService(c *configs.DbConfig, m *configs.EMailConfig) (*service.Service, error) {
	mailSender := email.NewMailSender(m)
	sqlDB, err := configs.NewDb(c)
	if err != nil {
		return nil, err
	}
	userRepo := dao.NewUserRepo(sqlDB)
	serviceService := service.NewService(mailSender, userRepo)
	return serviceService, nil
}
