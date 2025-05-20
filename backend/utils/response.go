package utils

// Standar response untuk API
type Response struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Error   string      `json:"error,omitempty"`
}

// SuccessResponse buat response sukses
func SuccessResponse(message string, data interface{}) Response {
	return Response{
		Status:  true,
		Message: message,
		Data:    data,
	}
}

// ErrorResponse buat response error
func ErrorResponse(message string, err error) Response {
	errorMessage := ""
	if err != nil {
		errorMessage = err.Error()
	}
	return Response{
		Status:  false,
		Message: message,
		Error:   errorMessage,
	}
}
