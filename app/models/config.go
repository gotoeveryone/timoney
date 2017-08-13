package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"

	"github.com/revel/revel"
)

type (
	// AppConfig 設定
	AppConfig struct {
		Redis redis `json:"redis"`
		DB    db    `json:"db"`
		Mail  mail  `json:"mail"`
	}

	// Redis接続設定
	redis struct {
		Host string `json:"host"`
		Port int    `json:"port"`
		Auth string `json:"auth"`
	}

	// データベース接続設定
	db struct {
		Name     string `json:"name"`
		Host     string `json:"host"`
		Port     int    `json:"port"`
		User     string `json:"user"`
		Password string `json:"password"`
	}

	// メール接続設定
	mail struct {
		SMTP      string   `json:"smtp"`
		Port      int      `json:"port"`
		User      string   `json:"user"`
		Password  string   `json:"password"`
		From      string   `json:"from"`
		FromAlias string   `json:"fromAlias"`
		To        []string `json:"to"`
	}
)

// LoadConfig アプリケーション設定をJSONファイルから読み込む
func LoadConfig(config *AppConfig) {
	jsonValue, err := ioutil.ReadFile(fmt.Sprintf("%s/config.json", revel.BasePath))
	if err != nil {
		revel.ERROR.Fatalln(err)
	}
	if err := json.Unmarshal(jsonValue, &config); err != nil {
		revel.ERROR.Fatalln(err)
	}
}
