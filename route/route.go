package route

import (
	"octagon/controller"

	"github.com/gofiber/fiber/v2"
)

func RouteInit(c *fiber.App) {
	c.Get("/", controller.GetHelloWorld)
	c.Get("/post/:id", controller.GetPost)
	c.Get("/posts", controller.GetPosts)
	c.Post("/post", controller.AddPost)
	c.Delete("/post", controller.DeletePost)
}
