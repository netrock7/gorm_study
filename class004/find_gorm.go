package main

import "fmt"

func TestFind() {
	//下面两行写法二选一
	var result = make(map[string]interface{})
	//result := map[string]interface{}
	GLOBAL_DB.First(&TestUser{}).First(&result)
	fmt.Println(result)

}
