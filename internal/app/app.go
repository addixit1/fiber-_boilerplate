package app

import (
	"github.com/addixit1/fiber-boilerplate/internal/config"
	errors "github.com/addixit1/fiber-boilerplate/internal/error"
	"github.com/addixit1/fiber-boilerplate/internal/lib/dbConnection"
	"github.com/addixit1/fiber-boilerplate/internal/lib/locale"
	"github.com/addixit1/fiber-boilerplate/internal/lib/redis"
	"github.com/addixit1/fiber-boilerplate/internal/lib/swagger"
	"github.com/addixit1/fiber-boilerplate/internal/utils"
	"github.com/gofiber/fiber/v2"
)

func New() *fiber.App {
	// Load configuration
	config.LoadEnv()

	// Load locale files
	if err := locale.Load(); err != nil {
		utils.LogError("Failed to load locale files: " + err.Error())
	} else {
		utils.LogSuccess("Locale files loaded successfully")
	}

	// Print startup banner
	utils.LogStartup("Fiber Boilerplate API", "1.0.0", "3010")

	// Connect to databases
	dbConnection.ConnectMongo()
	redis.Init()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: errors.Handler,
	})

	registerMiddlewares(app)
	registerRoutes(app)
	swagger.Register(app)

	utils.LogServer("Swagger UI available at http://localhost:3010/swagger/index.html")
	utils.LogSuccess("Application initialized successfully!")

	return app
}
