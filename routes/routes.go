package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/igor-koniukhov/fiber-go-auth/controllers"
)

func Setup(app *fiber.App)  {
	app.Post("/api/register",controllers.Register)
	app.Post("/api/login",controllers.Login)
	app.Get("/api/user", controllers.User)
	app.Post("/api/logout", controllers.Logout)
	app.Post("/api/forgot", controllers.Forgot)
	app.Post("/api/reset", controllers.Reset)
}


