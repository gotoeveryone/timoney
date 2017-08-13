package models

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var (
	// Db 接続用インスタンス
	Db *gorm.DB
)

// InitDB テーブル初期化
func InitDB() {
	var err error
	// dsn := "goute:kazuki11@/igo?charset=utf8&parseTime=True&loc=Local"
	dsn := fmt.Sprintf("%s:%s@%s/%s?charset=utf8&parseTime=True&loc=Local",
		"myuser", "mypassword", "tcp(myhost:myport)", "mydbname")

	Db, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	Db.LogMode(true)
	revel.INFO.Println("Connected to database")
}
