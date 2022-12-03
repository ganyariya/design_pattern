package cor

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMiddleware(t *testing.T) {
	doubleHandler := &DoubleRequestHandler{}

	loggerMiddleware := &LoggerMiddleware{}
	maintenanceMiddleware := &MaintenanceMiddleware{active: true}
	responseWrapMiddleware := &ResponseWrapMiddleware{prefix: "{", suffix: "}"}

	requestHandler := NewWrapperHandler(loggerMiddleware, NewWrapperHandler(maintenanceMiddleware, NewWrapperHandler(responseWrapMiddleware, doubleHandler)))
	request := &ServerRequest{body: "Hello, World!"}
	response := requestHandler.handle(request)
	assert.Equal(t, "Maintenance Now", response.GetBody())

	maintenanceMiddleware = &MaintenanceMiddleware{active: false}
	requestHandler = NewWrapperHandler(loggerMiddleware, NewWrapperHandler(maintenanceMiddleware, NewWrapperHandler(responseWrapMiddleware, doubleHandler)))
	request = &ServerRequest{body: "Hello, World!"}
	response = requestHandler.handle(request)
	assert.Equal(t, "{Hello, World!Hello, World!}", response.GetBody())
}
