package controllers

import "github.com/gofiber/fiber"

type Controllers interface {
	PostFile(c *fiber.Ctx)
	GetData(c *fiber.Ctx)
}
