package route

import (
	"octagon/controller"
	"octagon/middlewares"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	user := c.Group("auth")
	user.Post("/registerUser", controller.Register)
	user.Post("/loginUser", controller.Login)
	user.Get("/logout", controller.Logout)
	user.Delete("/user", controller.DeleteUser)

	v1 := c.Group("/post")
	v1.Use(middlewares.JWTMiddleware())
	v1.Get("/posts", controller.GetPosts)
	v1.Get("/", controller.GetHelloWorld)
	v1.Get("/post/:id", controller.GetPost)
	v1.Get("/posts", controller.GetPosts)
	v1.Post("/post", controller.AddPost)
	v1.Delete("/post", controller.DeletePost)
}
