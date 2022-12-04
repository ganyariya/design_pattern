package facade

import "github.com/ganyariya/design_pattern/ChainOfResponsibility/go_chain_of_responsibility/cor"

type AppFacade struct {
	Middlewares  []cor.MiddlewareInterface
	EntryHandler cor.RequestHandlerInterface
}

func NewAppFacade() *AppFacade {
	return &AppFacade{Middlewares: make([]cor.MiddlewareInterface, 0), EntryHandler: nil}
}

func (a *AppFacade) AddMiddleware(m cor.MiddlewareInterface) {
	a.Middlewares = append(a.Middlewares, m)
}

func (a *AppFacade) SetEntryHandler(e cor.RequestHandlerInterface) {
	a.EntryHandler = e
}

func (a *AppFacade) Run(request cor.ServerRequestInterface) cor.ResponseInterface {
	handler := a.EntryHandler
	for _, m := range a.Middlewares {
		handler = cor.NewWrapperHandler(m, handler)
	}

	return handler.Handle(request)
}
