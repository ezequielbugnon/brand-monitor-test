package controllers

import (
	"github.com/ezequiel-bugnon/brandmonitor/services"
	"github.com/gofiber/fiber"
)

type FiberController struct {
	service services.Service
}

func NewController(service services.Service) *FiberController {
	return &FiberController{
		service,
	}
}

func (f *FiberController) PostFile(c *fiber.Ctx) {
	c.SendString("dardo")
}

func (f *FiberController) GetData(c *fiber.Ctx) {
	c.SendString("dardo")
}
