package main

import "fmt"

func TestFind() {
	//下面两行写法二选一
	// var result = make(map[string]interface{})
	//result := map[string]interface{}
	var user TestUser
	// GLOBAL_DB.Model(&TestUser{}).First(&result)
	GLOBAL_DB.Model(&TestUser{}).First(&user)
	// fmt.Println(result)
	fmt.Println(user)
}
