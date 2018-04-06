package mockback

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResponseWriterMock struct {
	t *testing.T
}

func (m ResponseWriterMock) Header() http.Header {
	return nil
}

func (m *ResponseWriterMock) Write(b []byte) (int, error) {
	assert.Equal(m.t, "It works!", string(b))
	return 0, nil
}

func (m ResponseWriterMock) WriteHeader(statusCode int) {
}

func TestMock(t *testing.T) {
	var m = ResponseWriterMock{t: t}
	Mock(&m, nil)
}
