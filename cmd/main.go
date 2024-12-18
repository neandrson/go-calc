package main

import (
	"github.com/neandrson/go-calc/internal/application"
)

func main() {
	app := application.New()
	//app.Run()
	app.RunServer()
}
