package main

import (
	"fmt"
	"errors"
	"database/sql"
	"github.com/Eric-WangHaitao/Go-0712/Week02/dao"
)

func main() {
	defer dao.Db.Close()
	user, err := dao.QueryUserById("123456")
	if errors.Is(err, sql.ErrNoRows} {
		fmt.Printf("data not found, %v\n", err)
		return
    }

	if err != nil {
		// unknown error
		return
	}

	fmt.Println("query user : ", user)
}
