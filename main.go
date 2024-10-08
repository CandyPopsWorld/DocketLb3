package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	// Создаем новый роутер Gin
	r := gin.Default()

	// Определяем простой маршрут GET
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// Запуск сервера на порту 8080
	r.Run() // по умолчанию запускается на :8080
}
