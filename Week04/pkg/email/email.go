package email

import (
	"github.com/google/wire"
	"github.com/Eric-WangHaitao/Go-0712/Week04/configs"
	"log"
)

type EMail interface {
	Send()
}

type eMail struct {
}

func (e *eMail) Send() {
	log.Println("send email")
}

func NewMailSender(m *configs.EMailConfig) *eMail {
	return &eMail{}
}

var EMailSet = wire.NewSet(NewMailSender, wire.Bind(new(EMail), new(*EMail)))
