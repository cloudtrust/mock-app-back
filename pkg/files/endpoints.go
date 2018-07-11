package files

import (
	"context"
	"net/http"
	"strconv"

	"github.com/go-kit/kit/endpoint"
)

// Endpoints wraps a service behind a set of endpoints.
type Endpoints struct {
	ListSomeEndpoint endpoint.Endpoint
}

// MakeListSomeEndpoint makes the ListSomeEndpoint.
func MakeListSomeEndpoint(component Component) endpoint.Endpoint {
	return func(ctx context.Context, req interface{}) (interface{}, error) {
		var parameters = req.(*http.Request).URL.Query()
		// If the query has a first and a rows parameter...
		if firstArr, ok := parameters["first"]; ok {
			if rowsArr, ok := parameters["rows"]; ok {
				// We parse them. If one of them couldn't be parsed, we give up
				var first, erra = strconv.ParseInt(firstArr[0], 10, 32)
				if erra != nil {
					return nil, erra
				}
				var rows, errb = strconv.ParseInt(rowsArr[0], 10, 32)
				if errb != nil {
					return nil, errb
				}
				// We return a query of some files
				return component.ListSome(ctx, int32(first), int32(rows))
			}
		}
		// We return a query of all files
		return component.ListAll(ctx)
	}
}
