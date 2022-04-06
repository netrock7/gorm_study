package main

import "time"

type Model struct {
	UUID uint `gorm:"primaryKey""`
	//Time time.Time `gorm:"column:my_time"`
}

type TestUser struct {
	//gorm.Model
	Model Model `gorm:"embedded"`
	//ID           uint
	Name string `gorm:"default:zm"`
	//Email        *string
	Age uint8 `gorm:"comment:年龄"` // gorm的标签功能
	//Birthday     *time.Time
	//MemberNumber sql.NullString
	//ActivedAt    sql.NullTime
	CreatedAt time.Time
	UpdatedAt time.Time
}

func TestUserCreate() {
	GLOBAL_DB.AutoMigrate(&TestUser{})
}
