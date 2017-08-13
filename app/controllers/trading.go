package controllers

import (
	"time"
	"timoney/app/forms"
	"timoney/app/models"

	"github.com/revel/revel"
)

// Tradings 共通コントローラ
type Tradings struct {
	*revel.Controller
	models.Transactional
}

// API 初期表示
func (c Tradings) API() revel.Result {
	tradings := []models.Trading{}
	c.Txn.Preload("Account").Find(&tradings)
	return c.RenderJSON(tradings)
}

// Index 初期表示
func (c Tradings) Index() revel.Result {
	tradings := []models.Trading{}
	c.Txn.Preload("Account").Find(&tradings)

	form := forms.TradingForm{}
	return c.Render(form, tradings)
}

// Save 保存処理
func (c Tradings) Save(form forms.TradingForm) revel.Result {
	revel.INFO.Println(form)
	// バリデーションエラー
	form.Validate(c.Validation, c.Request.Locale)
	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Tradings.Index)
	}

	// 取引日の変換エラー
	traded, err := time.Parse("2006/1/2", form.Traded)
	if err != nil {
		c.Validation.Error("日付が不正です。").Key("tradings.form.Traded.invalid")
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(Tradings.Index)
	}

	trading := models.Trading{
		Name:      &form.Name,
		AccountID: form.AccountID,
		Traded:    &traded,
		Payment:   form.Payment,
	}

	if err := c.Txn.Create(&trading).Error; err != nil {
		return c.RenderError(err)
	}

	return c.Redirect(Tradings.Index)
}
