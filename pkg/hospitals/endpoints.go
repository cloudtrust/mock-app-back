package hospitals

import (
	"context"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps a service behind a set of endpoints.
type Endpoints struct {
	ListAllHospitalsEndpoint   endpoint.Endpoint
	ListAllDepartmentsEndpoint endpoint.Endpoint
	ListAllDoctorsEndpoint     endpoint.Endpoint
}

// MakeListAllHospitalsEndpoint makes the ListAllHospitalsEndpoint.
func MakeListAllHospitalsEndpoint(component Component) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return component.ListAllHospitals(ctx)
	}
}

// MakeListAllDepartmentsEndpoint makes the ListAllDepartmentsEndpoint.
func MakeListAllDepartmentsEndpoint(component Component) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return component.ListAllDepartments(ctx)
	}
}

// MakeListAllDoctorsEndpoint makes the ListAllDoctorsEndpoint.
func MakeListAllDoctorsEndpoint(component Component) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		return component.ListAllDoctors(ctx)
	}
}
