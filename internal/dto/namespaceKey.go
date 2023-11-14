package dto

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

type NamespaceKey struct {
	APIGroupKey

	Namespace string
}

func (nk *NamespaceKey) Fill(c *fiber.Ctx) error {
	err := nk.APIGroupKey.Fill(c)
	if err != nil {
		return err
	}

	nk.Namespace = c.Params("namespace", "") // namespace is optional for global resources

	return nil
}

func (nk *NamespaceKey) UniqueKey() string {
	return fmt.Sprintf("%s/%s", nk.APIGroupKey.UniqueKey(), nk.Namespace)
}
