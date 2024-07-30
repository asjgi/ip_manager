package http

import (
	"github.com/gofiber/fiber/v2"
	"ip_manager/usecases"
	"strconv"
)

type Handler struct {
	manager *usecases.IPManager
}

func NewHandler(manager *usecases.IPManager) *Handler {
	return &Handler{manager: manager}
}

func (h *Handler) CreateSubnet(c *fiber.Ctx) error {
	var req struct {
		CIDR     string `json:"cidr"`
		ParentID int    `json:"parent_id"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ipRange, err := h.manager.CreateSubnet(req.CIDR, req.ParentID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(ipRange)
}

func (h *Handler) DeleteSubnet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.manager.DeleteSubnet(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) AllocateSubnet(c *fiber.Ctx) error {
	var req struct {
		ParentID int `json:"parent_id"`
		CIDR     int `json:"cidr"`
	}
	if err := c.BodyParser(&req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	ipRange, err := h.manager.AllocateSubnet(req.ParentID, req.CIDR)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusCreated).JSON(ipRange)
}

func (h *Handler) ReleaseSubnet(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	if err := h.manager.ReleaseSubnet(id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.SendStatus(fiber.StatusNoContent)
}

func (h *Handler) CheckIPAllocated(c *fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	allocated, err := h.manager.CheckIPAllocated(id)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{"allocated": allocated})
}
