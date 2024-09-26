package database

import (
	"StudentServicePlatform/internal/global"
	"fmt"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Init()*gorm.DB{
	user :=global.Config.GetString("mysql.user")
    pass :=global.Config.GetString("mysql.pass")
    host :=global.Config.GetString("mysql.host")
    port :=global.Config.GetString("mysql.port")
    name :=global.Config.GetString("mysql.DBname")

    dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, name)
    db,err:=gorm.Open(mysql.Open(dsn),&gorm.Config{
        DisableForeignKeyConstraintWhenMigrating: true, // 禁用外键约束,方便测试
    })
	if err!=nil{
		log.Fatal("Database :", err)
	}

	err = autoMigrate(db)
	if err != nil {
		log.Fatal("Database migrate failed:", err)
	}

	DB = db
	return DB
}

