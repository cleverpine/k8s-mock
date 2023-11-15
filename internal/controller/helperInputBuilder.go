package controller

import "github.com/gofiber/fiber/v2"

func makeInputBuilder(c *fiber.Ctx) *inputBuilder {
	return &inputBuilder{
		c: c,
	}
}

type inputBuilder struct {
	c   *fiber.Ctx
	err error
}

func (ib *inputBuilder) InBody(v Validatable) *inputBuilder {
	return ib.
		setError(ib.c.BodyParser(v)).
		setError(v.Validate())
}

func (ib *inputBuilder) InURL(v Validatable) *inputBuilder {
	return ib.
		setError(ib.c.ParamsParser(v)).
		setError(v.Validate())
}

func (ib *inputBuilder) InQuery(v Validatable) *inputBuilder {
	return ib.
		setError(ib.c.QueryParser(v)).
		setError(v.Validate())
}

func (ib *inputBuilder) setError(err error) *inputBuilder {
	ib.err = err
	return ib
}

func (ib *inputBuilder) Error() error {
	return ib.err
}

type Validatable interface {
	Validate() error
}
