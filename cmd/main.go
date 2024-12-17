package main

import (
	"github.com/neandrson/go-calc/tree/main/internal/application"
)

func main() {
	app := application.New()
	app.RunServer()
}
