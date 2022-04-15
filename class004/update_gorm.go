package main

func TestUpdate() {
	// update	只更新选择的字段
	// updates	更新所有字段	此时有两种形式	一种为map，一种为机构体，结构体0值不参与更新
	// save		无论如何都更新	所有的内容	包括0值
	GLOBAL_DB.Model(&TestUser{}).Where("name = ?", "aaa").Update()
}
