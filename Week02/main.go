package main

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/Eric-WangHaitao/Go-0712/Week02/dao"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	defer dao.Db.Close()
	user, err := dao.QueryUserById("123456")
	if errors.Is(err, sql.ErrNoRows) {
		fmt.Printf("data not found, %v\n", err)
		return
	}

	if err != nil {
		// unknown error
		return
	}

	fmt.Println("query user : ", user)
}
