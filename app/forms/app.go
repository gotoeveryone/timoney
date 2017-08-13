package forms

import "github.com/revel/revel"

// LoginForm ログインフォーム
type LoginForm struct {
	Account  string
	Password string
}

// Validate バリデーション
func (c LoginForm) Validate(v *revel.Validation, locale string) {
	v.Check(c.Account, revel.ValidRequired(), revel.ValidMinSize(6))
	v.Check(c.Password, revel.ValidRequired(), revel.ValidMinSize(6))
}

// TradingForm 取引入力フォーム
type TradingForm struct {
	AccountID          int
	Traded             string
	Name               string
	Means              string
	PaymentDueDate     string
	Summary            string
	Suppliers          string
	Payment            int
	DistributionRatios int8
}

// Validate バリデーション
func (c TradingForm) Validate(v *revel.Validation, locale string) {
	v.Check(c.AccountID, revel.ValidRequired())
	v.Check(c.Traded, revel.ValidRequired(), revel.ValidMinSize(8), revel.ValidMaxSize(10))
	v.Check(c.Payment, revel.ValidRequired(), revel.ValidRange(1, 9999999))
}
