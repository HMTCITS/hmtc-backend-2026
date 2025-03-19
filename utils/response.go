package utils

type Response struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
	Error   any    `json:"error,omitempty"`
}

func ResponseSuccess(msg string, data any) Response {
	return Response{
		Status:  true,
		Message: msg,
		Data:    data,
	}
}

func ResponseFailed(msg string, err any) Response {
	return Response{
		Status:  false,
		Message: msg,
		Error:   err,
	}
}
