package cor

type ResponseInterface interface {
	GetBody() string
}

type Response struct {
	body string
}

func NewResponse(body string) *Response {
	return &Response{body}
}

func (r *Response) GetBody() string {
	return r.body
}
