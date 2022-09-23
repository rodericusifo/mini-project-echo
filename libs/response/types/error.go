package type_response

type ResponseError struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	Error   error  `json:"error,omitempty"`
}
