package main

import (
	"fmt"
)

// 单条查询
// func TestFind() {
// 	//下面两行写法二选一
// 	// var result = make(map[string]interface{})
// 	//result := map[string]interface{}
// 	var user TestUser
// 	// GLOBAL_DB.Model(&TestUser{}).First(&result)
// 	// db_req := GLOBAL_DB.Model(&TestUser{}).First(&user, 173)
// 	// db_req := GLOBAL_DB.Where("name = ?", "ccc").First(&user)
// 	db_req := GLOBAL_DB.Where(TestUser{Name: "ddd3"}).Or("name = ?", "eee").First(&user)
// 	// db_req := GLOBAL_DB.Where(map[string]interface{}{
// 	// 	"name": "bbb",
// 	// })
// 	// fmt.Println(result)
// 	fmt.Println(user)
// 	fmt.Println(db_req.Error)
// 	fmt.Println(errors.Is(db_req.Error, gorm.ErrRecordNotFound))
// }

// // 批量查询
// func TestFind() {
// 	var User []TestUser
// 	GLOBAL_DB.Where("name = ?", "aaa").Find(&User)
// 	fmt.Println(User)
// }

// 指定查询结果字段
type UserInfo struct {
	Name string
	Age  uint8
}

func TestFind() {
	var u []UserInfo
	GLOBAL_DB.Model(&TestUser{}).Where("name = ?", "aaa").Find(&u)
	fmt.Println(u)
}
