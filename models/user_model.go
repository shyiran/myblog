package models

import (
	"fmt"
	"myblog/common"
)

type User struct {
	Id         int
	Username   string
	Password   string
	Email      string
	Status     int
	Createtime int64
}

//按条件查询
func QueryUserWightCon(con string) int {
	sql := fmt.Sprintf("select id from users %s", con)
	fmt.Println(sql)
	row := common.QueryRowDB(sql)
	id := 1
	row.Scan(&id)
	return id
}
//插入
func InsertUser(user User) (int64, error) {
	return common.ModifyDB("insert into tb_user(username,password,email,status,createtime) values (?,?,?,?)",
		user.Username, user.Password,user.Email,user.Status, user.Createtime)
}
//根据用户名查询id
func QueryUserWithUsername(username string) int {
	sql := "where username=" + username
	return QueryUserWightCon(sql)
}

//根据用户名和密码，查询id
func QueryUserWithParam(username, password string) int {
	sql := fmt.Sprintf("where username='%s' and password='%s'", username, password)
	fmt.Println(sql+"AAAAAAAAAAAAAAaaaaa")
	return QueryUserWightCon(sql)
}
