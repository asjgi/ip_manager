package router

import (
	router "ip_manager/external/http"
	httpHandler "ip_manager/interfaces/http"
	usecases "ip_manager/usecases"
	"net/http/httptest"
	"testing"

	"github.com/gofiber/fiber/v2"
	"github.com/stretchr/testify/assert"
)

func setupTestRouter() *fiber.App {
	manager := usecases.NewIPManager()
	handler := httpHandler.NewHandler(manager)
	return router.NewRouter(handler)
}

func TestRouterCreateSubnet(t *testing.T) {
	app := setupTestRouter()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestRouterDeleteSubnet(t *testing.T) {
	app := setupTestRouter()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("DELETE", "/subnets/1", nil)
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 204, resp.StatusCode)
}

func TestRouterAllocateSubnet(t *testing.T) {
	app := setupTestRouter()

	req := httptest.NewRequest("POST", "/subnets", nil)
	req.Header.Set("Content-Type", "application/json")
	app.Test(req, -1)

	req = httptest.NewRequest("POST", "/subnets/allocate", nil)
	req.Header.Set("Content-Type", "application/json")
	resp, _ := app.Test(req, -1)

	assert.Equal(t, 201, resp.StatusCode)
}

func TestRouterReleaseSubnet(t *testing.T) {
	app := setupTestRouter()

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
