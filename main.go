package main

import (
	"github.com/sundayj1213/go-todo-rest-api/app"
	"github.com/sundayj1213/go-todo-rest-api/config"
)

func main() {
	config := config.GetConfig()

	app := &app.App{}
	app.Initialize(config)
	app.Run(":3000")
}
