package controllers

import (
	"timoney/app/forms"
	"timoney/app/models"

	"github.com/revel/revel"
)

// App 共通コントローラ
type App struct {
	*revel.Controller
	models.Transactional
}

// Index 初期表示
func (c App) Index() revel.Result {
	form := forms.LoginForm{}
	return c.Render(form)
}

// Login ログイン
func (c App) Login(form forms.LoginForm) revel.Result {
	revel.INFO.Println(form)
	form.Validate(c.Validation, c.Request.Locale)

	if c.Validation.HasErrors() {
		c.Validation.Keep()
		c.FlashParams()
		return c.Redirect(App.Index)
	}

	// ログイン成功
	return c.Redirect(Tradings.Index)
}
