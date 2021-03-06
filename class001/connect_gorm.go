package main

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

//func main() {
//	dsn := "go:ha9nD784yO7MFB@tcp(10.38.1.12:3306)/gorm_class?charset=utf8mb4&parseTime=True&loc=Local"
//	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
//	fmt.Println(db, err)
//}

func main() {
	db, _ := gorm.Open(mysql.New(mysql.Config{
		DSN:               "admin_wr:HjA86jqgm9mZ@tcp(localhost:3306)/gorm_class?charset=utf8&parseTime=True&loc=Local", // DSN data source name
		DefaultStringSize: 256,                                                                                          // string 类型字段的默认长度
	}), &gorm.Config{
		SkipDefaultTransaction: false,
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   "gva_", // 表名前缀，例如表名是user，这里就是t_user
			SingularTable: true,   //	使用单数表名，user的表名应该是t_users
		},
		DisableForeignKeyConstraintWhenMigrating: true, //	关闭建立物理外键
	})
	//	连接池配置信息
	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)           //	连接池中最大空闲连接数
	sqlDB.SetMaxOpenConns(100)          //	连接池最大连接数
	sqlDB.SetConnMaxLifetime(time.Hour) //	连接池中连接的最大可复用时间

	type User struct {
		Name string
	}
	//type User2 struct {
	//	Name string
	//}
	// 结构体自动创建表
	db.AutoMigrate(&User{})

	// 手工创建表
	//M := db.Migrator()
	//M.CreateTable(&User{})
	//fmt.Println(M.HasTable(&User{}))
	//M.DropTable(&User{})
	//fmt.Println(M.HasTable(&User{}))

	//if M.HasTable(&User{}) {
	//	err = M.RenameTable(&User{}, &User2{})
	//} else {
	//	err = M.RenameTable(&User2{}, &User{})
	//}
	//fmt.Println(db, err)
}
