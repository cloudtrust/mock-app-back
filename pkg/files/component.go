package files

import (
	"context"
)

// Component is the files business component interface.
type Component interface {
	ListAll(context.Context) ([]File, error)
}

// component is the files business component.
type component struct {
	module Module
}

func (c *component) ListAll(ctx context.Context) ([]File, error) {
	return c.module.ListAll(ctx)
}

// NewComponent returns a files business component
func NewComponent(module Module) Component {
	return &component{
		module: module,
	}
}
