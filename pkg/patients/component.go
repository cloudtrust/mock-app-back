package patients

import (
	"context"
)

// Component is the patient business component interface.
type Component interface {
	ListAll(context.Context) ([]Patient, error)
}

// component is the patient business component.
type component struct {
	module Module
}

func (c *component) ListAll(ctx context.Context) ([]Patient, error) {
	return c.module.ListAll(ctx)
}

// NewComponent returns a patient business component
func NewComponent(module Module) Component {
	return &component{
		module: module,
	}
}
