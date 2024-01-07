package middlewares

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func Cors(app *fiber.App) {
	
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://martinabelrr.github.io/GO_BLOGFRONTEND",		
		AllowHeaders:  "Origin, Content-Type, Accept",
		AllowCredentials: true,
	}))
}	