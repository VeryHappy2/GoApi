package response

type Response[T any] struct {
	Status string `json:"status"`
	Data   T      `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

type DefaultResponse struct {
	Status string `json:"status"`
	Data   int8   `json:"data,omitempty"`
	Error  string `json:"error,omitempty"`
}

const (
	StatusOK    = "OK"
	StatusError = "Error"
)

func OK[T any](data *T) Response[T] {
	return Response[T]{
		Status: StatusOK,
		Data:   *data,
	}
}

func Error[T any](msg string, data *T) Response[T] {
	return Response[T]{
		Status: StatusError,
		Error:  msg,
		Data:   *data,
	}
}
