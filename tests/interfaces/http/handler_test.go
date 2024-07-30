package http

import (
	httpHandler "ip_manager/interfaces/http"
	"ip_manager/usecases"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestApp() *fiber.App {
	manager := usecases.NewIPManager()
	handler := httpHandler.NewHandler(manager)
	app := fiber.New()

	app.Post("/subnets", handler.CreateSubnet)
	app.Delete("/subnets/:id", handler.DeleteSubnet)
	app.Post("/subnets/allocate", handler.AllocateSubnet)
	app.Post("/subnets/release/:id", handler.ReleaseSubnet)

	return app
}

func TestCreateSubnet(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestDeleteSubnet(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("DELETE", "/subnets/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestAllocateSubnet(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("POST", "/subnets/allocate", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestReleaseSubnet(t *testing.T) {
	app := setupTestApp()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("POST", "/subnets/allocate", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("POST", "/subnets/release/2", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 204, resp.StatusCode)
}
