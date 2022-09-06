package router

import (
	"server/handler"

	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

func NewEchoRouter() *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover()) //errorが起きてもpanicを起こさずrecoveryする

	e.POST("/users", handler.UserHandler{}.Create)
	e.GET("/users", handler.UserHandler{}.AllRead)
	e.GET("/users/:id", handler.UserHandler{}.Read)
	e.PUT("/users/:id", handler.UserHandler{}.Update)
	e.DELETE("/users/:id", handler.UserHandler{}.Delete)

	return e
}
