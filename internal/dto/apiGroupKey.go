package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type APIGroupKey struct {
	APIGroup string
	Version  string
}

func (agk *APIGroupKey) Fill(c *fiber.Ctx) error {
	agk.APIGroup = c.Params("apiGroup")
	agk.Version = c.Params("version")

	return nil
}

func (agk *APIGroupKey) UniqueKey() string {
	return fmt.Sprintf("%s/%s", agk.APIGroup, agk.Version)
}
