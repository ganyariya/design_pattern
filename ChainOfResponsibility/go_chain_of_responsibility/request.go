package cor

type ServerRequestInterface interface {
	GetMethod() string
	GetBody() string
}

type ServerRequest struct {
	method string
	body   string
}

func NewServerRequest(method, body string) *ServerRequest {
	return &ServerRequest{method, body}
}

func (s *ServerRequest) GetMethod() string {
	return s.method
}

func (s *ServerRequest) GetBody() string {
	return s.body
}
