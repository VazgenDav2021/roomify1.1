package main

import (
	"roomify-backend/config"
	"roomify-backend/models"
	"roomify-backend/routes"

	"github.com/labstack/echo/v4"
)

func main() {
	// Подключаемся к базе данных
	config.ConnectDatabase()

	// Автоматическая миграция базы данных
	models.Migrate(config.DB)

	// Создаем новый сервер Echo
	e := echo.New()

	// Группа маршрутов с префиксом /api/v1
	api := e.Group("/api/v1")

	// Регистрируем роуты
	routes.UserRoutes(api)  // Используем эту группу маршрутов

	// Запуск сервера
	e.Logger.Fatal(e.Start(":8080"))
}
