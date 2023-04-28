package main

import (
	"log"

	"github.com/CyberSleeper/backend-oprec-ristek/configs"
	"github.com/CyberSleeper/backend-oprec-ristek/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
)

func init() {
	config, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatalln("Failed to load environment variables!\n", err.Error())
	}
	configs.ConnectDB(&config)
}

func main() {
	app := fiber.New()
	micro := fiber.New()

	app.Mount("/", micro)
	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "http://localhost:3000",
		AllowHeaders:     "Origin, Content-Type, Accept",
		AllowMethods:     "GET, POST, PATCH, DELETE",
		AllowCredentials: true,
	}))

	micro.Route("/posts", func(router fiber.Router) {
		router.Post("/", controllers.CreatePostHandler)
		router.Get("", controllers.GetPostsHandler)
	})
	micro.Route("/posts/:postId", func(router fiber.Router) {
		router.Delete("", controllers.DeletePostHandler)
		router.Get("", controllers.GetPostByIdHandler)
		router.Patch("", controllers.UpdatePostHandler)
	})
	micro.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{
			"status":  "success",
			"message": "Success",
		})
	})

	log.Fatal(app.Listen(":8000"))
}
