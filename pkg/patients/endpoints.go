package patients

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps a service behind a set of endpoints.
type Endpoints struct {
	ListAllPatientsEndpoint endpoint.Endpoint
}

// MakeListAllPatientsEndpoint makes the ListAllPatientsEndpoint.
func MakeListAllPatientsEndpoint(component Component) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return component.ListAll(ctx)
	}
}
