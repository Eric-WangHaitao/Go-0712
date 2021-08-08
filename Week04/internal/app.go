package internal

import (
	"github.com/Eric-WangHaitao/Go-0712/Week04/configs"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/di"
)

//对外服务应用
type app struct {
}

func NewApp() *app {
	return &app{}
}

func (a *app) Run() error {
	// db配置
	dbConfig := &configs.DbConfig{}
	// 邮件配置
	mailConfig := &configs.EMailConfig{}
	s, err := di.NewService(dbConfig, mailConfig)
	if err != nil {
		return err
	}
	s.UserSignUp()
	return nil
}