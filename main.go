package main

import (
	"github.com/Onelvay/HL-architecture/internal/app"
)

// @title Course Service API
// @version 1.0
// @description a course service api in golang

// @host localhost:8080
// @BasePath /

// securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization

func main() {
	app.Run()
}
