import (
"fmt"
"gorm.io/driver/mysql"
"gorm.io/gorm"
"gorm.io/gorm/schema"
"time"
)
func main() {
db, err := gorm.Open(mysql.New(mysql.Config{
DSN:               "go:ha9nD784yO7MFB@tcp(10.38.1.12:3306)/gorm_class?charset=utf8&parseTime=True&loc=Local", // DSN data source name
DefaultStringSize: 256,                                                                                       // string 类型字段的默认长度
}), &gorm.Config{
SkipDefaultTransaction: false,
NamingStrategy: schema.NamingStrategy{
TablePrefix:   "gva_", // 表名前缀，例如表名是user，这里就是t_user
SingularTable: true,   //	使用单数表名，user的表名应该是t_users
},
DisableForeignKeyConstraintWhenMigrating: true, //	关闭建立物理外键
})
//	连接池配置信息
sqlDB, err := db.DB()
sqlDB.SetMaxIdleConns(10)           //	连接池中最大空闲连接数
sqlDB.SetMaxOpenConns(100)          //	连接池最大连接数
sqlDB.SetConnMaxLifetime(time.Hour) //	连接池中连接的最大可复用时间

	fmt.Println(db, err)
}
