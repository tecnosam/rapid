package http

type HTTPAdapter interface {
	Invoke(request HTTPRequest) HTTPResponse
}
