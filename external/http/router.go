package http

import (
	"github.com/gofiber/fiber/v2"
	"ip_manager/interfaces/http"
)

func NewRouter(handler *http.Handler) *fiber.App {
	app := fiber.New()

	app.Post("/subnets", handler.CreateSubnet)
	app.Delete("/subnets/:id", handler.DeleteSubnet)
	app.Post("/subnets/allocate", handler.AllocateSubnet)
	app.Post("/subnets/release/:id", handler.ReleaseSubnet)
	app.Get("/subnets/:id/allocated", handler.CheckIPAllocated)

	return app
}

func StartServer(handler *http.Handler, port string) {
	app := NewRouter(handler)
	app.Listen(":" + port)
}
