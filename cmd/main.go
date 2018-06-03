package main

import (
	"shorten/app"
)

func main() {
	shortenApp := app.NewApp()
	if shortenApp == nil {
		panic("shorten url service app is nil")
	}
	if err := shortenApp.Start(); err != nil {
		panic("app start error:" + err.Error())
	}
}
