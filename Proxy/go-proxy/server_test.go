package goproxy

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestServer(t *testing.T) {
	serverApp := &Server{}
	nginx := Nginx{serverApp: serverApp}

	assert.Equal(t, "admin request", nginx.HttpRequest("admin"))
	assert.Equal(t, "Hello! gana", nginx.HttpRequest("gana"))
}
