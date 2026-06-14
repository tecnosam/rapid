package http

type HTTPRequest struct {
	URL     string
	Method  string
	Headers *map[string]string
	Body    []byte
}

type HTTPResponse struct {
	Headers        *map[string]string
	HTTPStatusCode int
	Data           []byte
}

func (resp *HTTPResponse) IsSuccess() bool {
	return resp.HTTPStatusCode >= 200 && resp.HTTPStatusCode < 400
}
