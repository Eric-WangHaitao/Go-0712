package dao

import (
	"database/sql"
	"github.com/pkg/errors"
)

type User struct {
	Id   string
	Name string
}

var Db *sql.DB

func init() {
	var err error
	Db, err = sql.Open("mysql", "root:123456@tcp(127.0.0.1:3306)/test")
	if err != nil {
		panic(err)
	}
}

func QueryUserById(id string) (User, error) {
	var user User
	row := Db.QueryRow("select id ,name from user where id = ?", id)
	err := row.Scan(&user.Id, &user.Name)
	if err != nil {
		return user, errors.Wrap(err, "QueryUserById err")
	}
	return user, nil
}
