package routes

import (
	"github.com/gofiber/fiber/v2"

	"github.com/fprasty/GoApiWijaya/middleware"

	"github.com/fprasty/GoApiWijaya/controllers"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middleware.IsAuthenticate)
	//Post
	app.Post("api/post", controllers.CreatePost)
	app.Get("/api/allpost", controllers.AllPost)
	app.Get("/api/allpost/:id", controllers.DetailPost)
	app.Put("/api/updatepost/:id", controllers.UpdatePost)
	app.Get("/api/uniquepost", controllers.UniquePost)
	app.Delete("/api/deletepost/:id", controllers.DeletePost)
	//Comment
	app.Post("api/comment", controllers.CreateComment)
	app.Get("/api/allcomment", controllers.AllComment)
	app.Get("/api/allcomment/:id", controllers.DetailComment)
	app.Put("/api/updatecomment/:id", controllers.UpdateComment)
	app.Get("/api/uniquecomment", controllers.UniqueComment)
	app.Delete("/api/deletecomment/:id", controllers.DeleteComment)

}
