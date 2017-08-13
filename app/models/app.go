package models

import "time"

// Base 基底モデル
type Base struct {
	ID uint `gorm:"primary_key" json:"id"`
}

// TimestampBase タイムスタンプ付き基底モデル
type TimestampBase struct {
	Base
	CreatedAt time.Time `gorm:"column:created;type:datetime"`
	UpdatedAt time.Time `gorm:"column:modified;type:datetime"`
}

// Account 科目
type Account struct {
	Base
	Name string `json:"name"`
}

// Trading 取引
type Trading struct {
	TimestampBase
	AccountID          int        `json:"-"`
	Account            *Account   `gorm:"ForeignKey:AccountID" json:"account"`
	Traded             *time.Time `json:"traded;type:date"`
	Name               *string    `json:"name"`
	Means              *string    `json:"means"`
	PaymentDueDate     *time.Time `json:"paymentDueDate;type:date"`
	Summary            *string    `json:"summary"`
	Suppliers          *string    `json:"suppliers"`
	Payment            int        `json:"payment"`
	DistributionRatios *int8      `json:"distributionRatio"`
}

// TradingMean 取引手段
type TradingMean struct {
	Base
	Name string `json:"name"`
}

// FavoriteTrading よく使う取引
type FavoriteTrading struct {
	TimestampBase
	Keyword            string
	AccountID          int `json:"-"`
	TradingMeanID      int `json:"-"`
	Summary            *string
	Suppliers          *string
	Payment            *int
	DistributionRatios *int8
	Created            *time.Time
	Modified           *time.Time

	Account     Account     `gorm:"ForeignKey:AccountID" json:"account"`
	TradingMean TradingMean `gorm:"ForeignKey:TradingMeanId" json:"trading"`
}
