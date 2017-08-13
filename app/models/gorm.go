package models

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	"github.com/revel/revel"
)

var (
	// データベース接続用インスタンス
	dbManager *gorm.DB
)

// InitDB テーブル初期化
func InitDB() {
	// 設定ファイル読み出し
	config := AppConfig{}
	LoadConfig(&config)
	revel.INFO.Println(config)

	var err error
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=%s",
		config.DB.User,
		config.DB.Password,
		config.DB.Host,
		config.DB.Port,
		config.DB.Name,
		"Asia%2FTokyo",
	)

	dbManager, err = gorm.Open("mysql", dsn)
	if err != nil {
		panic(err)
	}
	dbManager.LogMode(true)
	revel.INFO.Println("Connected to database")
}

// Transactional トランザクション
type Transactional struct {
	*revel.Controller
	Txn *gorm.DB
}

// Begin トランザクション開始
func (c *Transactional) Begin() revel.Result {
	revel.TRACE.Println("トランザクションを開始します。")
	tx := dbManager.Begin()
	if err := tx.Error; err != nil {
		panic(err)
	}
	c.Txn = tx
	return nil
}

// Commit トランザクションのコミット
func (c *Transactional) Commit() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Commit()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	revel.TRACE.Printf("トランザクションをコミットしました。！")
	c.Txn = nil
	return nil
}

// Rollback トランザクションのロールバック
func (c *Transactional) Rollback() revel.Result {
	if c.Txn == nil {
		return nil
	}
	c.Txn.Rollback()
	if err := c.Txn.Error; err != nil && err != sql.ErrTxDone {
		panic(err)
	}
	revel.ERROR.Printf("トランザクションをロールバックしました。")
	c.Txn = nil
	return nil
}
