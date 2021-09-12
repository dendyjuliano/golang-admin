package routes

import (
	"go-admin/controllers"
	"go-admin/middlewares"

	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controllers.Register)
	app.Post("/api/login", controllers.Login)

	app.Use(middlewares.IsAuthenticated)

	app.Put("/api/user/info", controllers.UpdateInfo)
	app.Put("/api/user/password", controllers.UpdatePassword)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)

	app.Get("/api/users", controllers.AllUsers)
	app.Post("/api/users", controllers.CreateUser)
	app.Get("/api/users/:id", controllers.GetUser)
	app.Put("/api/users/:id", controllers.UpdateUser)
	app.Delete("/api/users/:id", controllers.DeleteUser)

	app.Get("/api/roles", controllers.AllRole)
	app.Post("/api/roles", controllers.CreateRole)
	app.Get("/api/roles/:id", controllers.GetRole)
	app.Put("/api/roles/:id", controllers.UpdateRole)
	app.Delete("/api/roles/:id", controllers.DeleteRole)

	app.Get("/api/product", controllers.AllProduct)
	app.Post("/api/product", controllers.CreateProduct)
	app.Get("/api/product/:id", controllers.GetProduct)
	app.Put("/api/product/:id", controllers.UpdateProduct)
	app.Delete("/api/product/:id", controllers.DeleteProduct)

	app.Get("/api/permissions", controllers.AllPermission)

	app.Post("/api/upload", controllers.Upload)
	app.Static("/api/uploads", "./uploads")

	app.Get("/api/order", controllers.AllOrder)
	app.Post("/api/export", controllers.Export)
	app.Get("/api/chart", controllers.Chart)
}
