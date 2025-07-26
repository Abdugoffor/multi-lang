package admin_cmd

import (
	"log"
	language_handler "project/module/admin/language/handler"
	post_handler "project/module/admin/post/handler"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Cmd(route *echo.Echo, db *gorm.DB, log *log.Logger) {

	routerGroup := route.Group("/api/v1/admin")
	{
		language_handler.NewLanguageHandler(routerGroup, db, log)
		post_handler.NewPostHandler(routerGroup, db, log)
	}
}
