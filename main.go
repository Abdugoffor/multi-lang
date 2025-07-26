package main

import (
	"log"
	"project/config"
	"project/helper"
	admin_cmd "project/module/admin"

	"github.com/labstack/echo/v4"
)

func main() {
	helper.LoadEnv()

	config.DBConnect()

	route := echo.New()

	admin_cmd.Cmd(route, config.DB, log.Default())

	route.Logger.Fatal(route.Start(":" + helper.ENV("HTTP_PORT")))
}
