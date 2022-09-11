package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"majoo.id/satu/config"
	"majoo.id/satu/controllers"
	"majoo.id/satu/middleware"
)

// Setup gofiber
func SetupRouters() *fiber.App {
	// Create new gofiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			code := fiber.StatusInternalServerError
			if e, ok := err.(*fiber.Error); ok {
				// Override status code if fiber.Error type
				code = e.Code
			}
			return c.Status(code).JSON(fiber.Map{"error": err.Error()})
		},
		AppName: config.APPLICATION_NAME,
	})

	// Set context default value
	app.Use(middleware.Setup)

	// Built in Middleware
	app.Use(recover.New(recover.Config{EnableStackTrace: true}))
	app.Use(compress.New(compress.Config{
		Level: compress.LevelBestSpeed,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusOK).JSON(fiber.Map{"error": false, "message": "API Uptime"})
	})

	// A. Fungsi Login
	app.Post("/login", controllers.Login)

	// B. Authorization
	app.Use(middleware.Authorization)

	// C. Laporan nama merchant, omzet
	app.Get("/opsi-c", controllers.OpsiC)

	// D. Laporan nama merchant, nama outlet, omzet
	app.Get("/opsi-d", controllers.OpsiD)

	// 404 routes
	app.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "Not found"})
	})

	return app
}
