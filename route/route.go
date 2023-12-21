package route

import (
	"octagon/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	v1 := c.Group("/post")
	v1.Get("/posts", controller.GetPosts)
	v1.Get("/", controller.GetHelloWorld)
	v1.Get("/post/:id", controller.GetPost)
	v1.Get("/posts", controller.GetPosts)
	v1.Post("/post", controller.AddPost)
	v1.Delete("/post", controller.DeletePost)

	user := c.Group("user")
	user.Get("/users", controller.GetPersons)
}
