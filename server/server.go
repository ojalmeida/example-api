package server

import (
	"example-api/config"
	"example-api/log"
	"example-api/server/routes"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/sirupsen/logrus"
)

var (
	app *fiber.App
)

func Init() {

	app = fiber.New(fiber.Config{
		ServerHeader:          "example-api",
		DisableStartupMessage: true,
	})

	app.Use(recover.New(recover.Config{
		EnableStackTrace: true,
	}))

	app.Use(func (c *fiber.Ctx) error  {
		log.Logger.WithFields(logrus.Fields{
			"url": c.OriginalURL(),
			"method": string(c.Context().Method()),
			"headers": fmt.Sprintf("%+v", c.GetReqHeaders()),
			"body": string(c.Body()),
		}).Debugln("request received")


		return c.Next()
	})

	routes.SetupApp(app)

}

func Run() (err error) {

	addr := fmt.Sprintf("%s:%d", config.Config.Server.Address, config.Config.Server.Port)
	err = app.Listen(addr)
	if err != nil {
		err = fmt.Errorf("unable to serve into %s:%d: %w", config.Config.Server.Address, config.Config.Server.Port, err)
	}

	return
}
