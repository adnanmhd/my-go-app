package controller

import (
	common "my-go-app/common"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	apmecho "go.elastic.co/apm/module/apmechov4"
)

func AssignRouting(e *echo.Echo) {
	e.Use(apmecho.Middleware()) // use apm
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "PUT", "DELETE"},
		AllowHeaders: []string{"*"},
	}))

	//assign middlewares
	e.Use(middleware.RequestID())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	//api handlers
	v1 := e.Group(common.Config.RootURL + "/v1" + common.Config.RootAPI)

	// Routes
	// v1.GET("/ping", ctl.Ping)
	v1.GET("/menus", GetMenus)
	// v1.GET("/menus/:id", getMenuById)
	v1.POST("/menus", AddMenu)
	// v1.DELETE("/menus/:id", deleteMenu)
	// v1.POST("/transactions", addTransaction)
	// v1.GET("/transactions", getTransactions)

}
