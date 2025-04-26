package routes

import (
	"roomify-backend/controllers"

	"github.com/labstack/echo/v4"
)

// Роуты для пользователей
func UserRoutes(e *echo.Group) {  // Принимаем группу, а не полный echo.Echo
	e.POST("/register", controllers.RegisterUser)
	e.GET("/account/:id", controllers.GetAccount)
	e.POST("/login", controllers.LoginUser)
	e.DELETE("/user/:id", controllers.DeleteUser)
	e.PUT("/user/:id", controllers.UpdateUser)
}
