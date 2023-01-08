package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/fprasty/GoApiWijaya/middleware"

	"github.com/fprasty/GoApiWijaya/controllers"
)

func Setup(app *fiber.App) {

	// Or extend your config for customization
	app.Use(cors.New(cors.Config{
		//AllowOrigins: "*",
		AllowMethods: "GET,POST,HEAD,OPTIONS,PUT,DELETE,PATCH",
		//AllowHeaders: "Origin, Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization,X-Requested-With",
		//ExposeHeaders:    "Origin",
		AllowCredentials: true,
	}))
	//Public
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	//app.Use(middleware.IsAuthenticate)
	//User
	app.Get("/api/getuser", middleware.IsAuthenticateUser, controllers.UserGetme)
	app.Get("/api/logout", middleware.IsAuthenticateUser, controllers.Logout)
	//app.Get("/api/alluser", middleware.IsAuthenticateUser, controllers.AllUser)
	app.Put("/api/updateuser/:id", middleware.IsAuthenticateUser, controllers.UpdateUser)
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
	//EndUser

	//Admin
	app.Post("/api/admin/register", controllers.AdminRegister)
	app.Post("/api/admin/login", controllers.AdminLogin)
	//User
	app.Get("/api/admin/getme", middleware.IsAuthenticateAdmin, controllers.AdminGetme)
	app.Get("/api/getuser/:id", middleware.IsAuthenticateAdmin, controllers.GetUser)
	app.Get("/api/getuser/uuid/:uuid", middleware.IsAuthenticateAdmin, controllers.GetUserUUID)
	app.Get("/api/admin/alluser", middleware.IsAuthenticateAdmin, controllers.AllUser)
}
