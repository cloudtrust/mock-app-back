package files

import (
	"context"
)

// Component is the files business component interface.
type Component interface {
	ListAll(context.Context) ([]File, error)
	ListSome(context.Context, int32, int32) ([]File, error)
	Count(context.Context) (int32, error)
}

// component is the files business component.
type component struct {
	module Module
}

func (c *component) ListAll(ctx context.Context) ([]File, error) {
	return c.module.ListAll(ctx)
}

func (c *component) ListSome(ctx context.Context, first int32, rows int32) ([]File, error) {
	return c.module.ListSome(ctx, first, rows)
}

func (c *component) Count(ctx context.Context) (int32, error) {
	return c.module.Count(ctx)
}

// NewComponent returns a files business component
func NewComponent(module Module) Component {
	return &component{
		module: module,
	}
}
