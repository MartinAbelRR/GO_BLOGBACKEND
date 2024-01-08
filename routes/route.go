package routes

import (
	"github.com/MartinAbelRR/blogbackend/controllers"
	// "github.com/MartinAbelRR/blogbackend/middlewares"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App){

	
	app.Post("/api/register", controllers.Resgister)
	app.Post("/api/login", controllers.Login)
	app.Post("/api/logout", controllers.Logout)
	
	// app.Use(middlewares.IsAuthenticate)
	
	app.Post("/api/post", controllers.CreatePost)
	app.Get("/api/allpost", controllers.AllPost)
	app.Get("/api/allpost/:id", controllers.DetailPost)
	app.Put("/api/updatepost/:id", controllers.UpdatePost)
	app.Get("/api/uniquepost", controllers.UniquePost)
	app.Delete("/api/deletepost/:id", controllers.DeletePost)
	app.Post("/api/upload-image", controllers.Upload)
	app.Static("/api/uploads", "./uploads")
}