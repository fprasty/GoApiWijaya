package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	//"github.com/fprasty/GoApiWijaya/middleware"

	"github.com/fprasty/GoApiWijaya/controllers"
)

func Setup(app *fiber.App) {
	// Default config
	//app.Use(cors.New())

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		AllowHeaders:     "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		ExposeHeaders:    "Origin",
		AllowCredentials: true,
	}))
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//app.Use(middleware.IsAuthenticate)

	app.Get("/api/logout", controllers.Logout)
	app.Get("/api/alluser", controllers.AllUser)
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
