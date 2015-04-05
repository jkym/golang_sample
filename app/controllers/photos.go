package controllers

import (
	"github.com/revel/revel"
)

type Photos struct {
	//*revel.Controller
	App
}

func (c Photos) Index() revel.Result {
	message := "photo index"
	return c.Render(message)
}

func (c Photos) Save() revel.Result {
	return c.Redirect(Photos.Index)
}
