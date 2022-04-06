package main

import "fmt"

func CreatedTest() {
	//db_req := GLOBAL_DB.Create(&TestUser{Name: "solo", Age: 24})
	//fmt.Println(db_req.Error, db_req.RowsAffected)

	db_req := GLOBAL_DB.Create(&[]TestUser{
		{Name: "aaa", Age: 21},
		{Name: "bbb", Age: 22},
		{Name: "ccc", Age: 23},
		{Name: "ddd", Age: 24},
		{Name: "eee", Age: 25},
	})

	fmt.Println(db_req.Error, db_req.RowsAffected)

	if db_req.Error != nil {
		fmt.Println("创建失败")
	} else {
		fmt.Println("创建成功")
	}
}
