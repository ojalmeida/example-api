package routes

import (
	"example-api/server/controllers"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"

	_ "example-api/docs"
)

func SetupApp(app *fiber.App) {

	apiRouter := app.Group("/api").Name(".api")
	v1Router := apiRouter.Group("/v1").Name(".v1")
	
	app.Get("/docs/*", swagger.HandlerDefault).Name(".docs")

	controllers.SetupRoutes(v1Router)

}
