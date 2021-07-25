package main

import (
   "fmt"
   "github.com/Eric-WangHaitao/Go-0712/dao"
)

func main() {
   defer dao.Db.Close()
   user, err := dao.QueryCustomerById("123456")
   if err != nil{
     fmt.Printf("query user err : %+v",err)
     return
   }
   fmt.Println("query user : ",user)
}
