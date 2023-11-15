package controller

import "github.com/gofiber/fiber/v2"

func NewMetadataController() *Metadata {
	return &Metadata{}
}

type Metadata struct {
}

func (ctrl *Metadata) Version(c *fiber.Ctx) error {
	return sendFile(c, "files/version.json")
}
