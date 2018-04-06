package mockback

import (
	"net/http"
	"testing"

	"github.com/stretchr/testify/assert"
)

type ResponseWriterMock struct {
	t        *testing.T
	expected string
}

func (m ResponseWriterMock) Header() http.Header {
	return nil
}

func (m *ResponseWriterMock) Write(b []byte) (int, error) {
	assert.Equal(m.t, m.expected, string(b))
	return 0, nil
}

func (m ResponseWriterMock) WriteHeader(statusCode int) {
}

func TestRoot(t *testing.T) {
	var m = ResponseWriterMock{t: t, expected: "It works!"}
	Root(&m, nil)
}

func TestDummyData(t *testing.T) {
	assert.True(t, len(GetDummyData()) == 2)
}
