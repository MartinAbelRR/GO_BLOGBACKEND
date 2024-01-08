package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) {
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: " https://martinabelrr.github.io, http://localhost:5173",		// https://martinabelrr.github.io
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
}	