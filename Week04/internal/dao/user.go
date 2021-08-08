package dao

import (
	"database/sql"
	"github.com/google/wire"
	"github.com/Eric-WangHaitao/Go-0712/Week04/internal/model"
	"log"
)

type UserRepository interface {
	AddUser()
}

type userRepo struct {
	*sql.DB
}

func (u *userRepo) AddUser() {
	user := &model.User{}
	user.Id = 1
	user.Name = "xiaoming"
	log.Println("add user :" + user.Name)
}

func NewUserRepo(db *sql.DB) *userRepo {
	return &userRepo{}
}

var UserSet = wire.NewSet(NewUserRepo, wire.Bind(new(UserRepository), new(*userRepo)))
