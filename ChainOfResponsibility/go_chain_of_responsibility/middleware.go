package cor

import (
	"log"
)

type MiddlewareInterface interface {
	process(ServerRequestInterface, RequestHandlerInterface) ResponseInterface
}

type LoggerMiddleware struct {
}

func (l *LoggerMiddleware) process(s ServerRequestInterface, r RequestHandlerInterface) ResponseInterface {
	log.Printf("Request Body: %s", s.GetBody())
	res := r.Handle(s)
	log.Printf("Response Body: %s", res.GetBody())
	return res
}

type MaintenanceMiddleware struct {
	active bool
}

func (m *MaintenanceMiddleware) process(s ServerRequestInterface, r RequestHandlerInterface) ResponseInterface {
	if m.active {
		return NewResponse("Maintenance Now")
	}
	return r.Handle(s)
}

type ResponseWrapMiddleware struct {
	prefix string
	suffix string
}

func NewResponseWrapMiddleware(prefix, suffix string) *ResponseWrapMiddleware {
	return &ResponseWrapMiddleware{prefix: prefix, suffix: suffix}
}

func (rwm *ResponseWrapMiddleware) process(s ServerRequestInterface, r RequestHandlerInterface) ResponseInterface {
	res := r.Handle(s)
	body := res.GetBody()
	return NewResponse(rwm.prefix + body + rwm.suffix)
}
