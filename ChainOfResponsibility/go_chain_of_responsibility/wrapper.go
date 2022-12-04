package cor

type WrapperHandler struct {
	middleware     MiddlewareInterface
	requestHandler RequestHandlerInterface
}

func NewWrapperHandler(middleware MiddlewareInterface, requestHandler RequestHandlerInterface) RequestHandlerInterface {
	return &WrapperHandler{middleware: middleware, requestHandler: requestHandler}
}

func (w *WrapperHandler) Handle(s ServerRequestInterface) ResponseInterface {
	return w.middleware.process(s, w.requestHandler)
}
