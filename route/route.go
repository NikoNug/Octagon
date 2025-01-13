package route

import (
	"octagon/controller"
	"octagon/middlewares"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
)

func RouteInit(c *fiber.App) {
	user := c.Group("auth")
	user.Post("/registerUser", controller.Register)
	user.Post("/loginUser", controller.Login)
	user.Get("/logout", controller.Logout)
	user.Delete("/user", controller.DeleteUser)

	protected := c.Group("/post")
	protected.Use(middlewares.JWTMiddleware())
	protected.Get("/posts", controller.GetPosts)
	protected.Get("/post/:id", controller.GetPost)
	protected.Post("/post", controller.AddPost)
	protected.Delete("/post", controller.DeletePost)

	protected.Get("/", controller.GetHelloWorld)

	c.Get("/login.html", func(c *fiber.Ctx) error {
		return c.Render("login", nil) // Render login template
	})

	chat := c.Group("/chat")
	chat.Use(middlewares.JWTMiddleware())
	chat.Get("/ws", controller.WebSocketHandler, websocket.New(controller.HandleWebSocket))

	chat.Get("/chat.html", func(c *fiber.Ctx) error {
		return c.Render("chat", fiber.Map{})
	})

}
