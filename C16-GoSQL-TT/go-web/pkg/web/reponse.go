package web

type Response struct {
	Code  int         `json:"status"`
	Data  interface{} `json:"data,omitempty"`
	Error string      `json:"error,omitempty"`
}

func NewReponse(code int, data interface{}, err error) Response {
	if code < 300 {
		return Response{code, data, ""}
	}
	return Response{code, nil, err.Error()}
}
