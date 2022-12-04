package facade

import (
	"testing"

	"github.com/ganyariya/design_pattern/ChainOfResponsibility/go_chain_of_responsibility/cor"
	"github.com/stretchr/testify/assert"
)

func TestFacade(t *testing.T) {
	app := NewAppFacade()
	app.SetEntryHandler(&cor.DoubleRequestHandler{})
	app.AddMiddleware(&cor.LoggerMiddleware{})
	app.AddMiddleware(&cor.MaintenanceMiddleware{})
	app.AddMiddleware(cor.NewResponseWrapMiddleware("{", "}"))

	resopnse := app.Run(cor.NewServerRequest("GET", "Hello"))

	assert.Equal(t, "{HelloHello}", resopnse.GetBody())
}
