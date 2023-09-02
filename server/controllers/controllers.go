package controllers

import (
	"github.com/gofiber/fiber/v2"
)

func SetupRoutes(router fiber.Router) {

	// ======clients controller ===========
	clientsRouter := router.Group("/clients")

	// create
	clientsRouter.Post("/", createClient)

	// read
	clientsRouter.Get("/", getClients)
	clientsRouter.Get("/:id", getClient)

	// update
	clientsRouter.Put("/:id", replaceClient)
	clientsRouter.Patch("/:id", updateClient)

	// delete
	clientsRouter.Delete("/:id", deleteClient)

	// ======================================
}
