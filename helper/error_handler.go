package helper

type Response struct {
	Code     int         `json:"code"`
	Status   bool        `json:"status"`
	Messagge string      `json:"message"`
	Data     interface{} `json:"data,omitempty"`
}

type ErrorResponse struct {
	Code     int         `json:"code"`
	Status   bool        `json:"status"`
	Messagge string      `json:"message"`
	Err      interface{} `json:"err,omitempty"`
}

func (r *Response) Error() string {
	return r.Messagge
}

func (r *ErrorResponse) Error() string {
	return r.Messagge
}

func NewSuccess(message string) error {
	return &Response{
		Code:     200,
		Status:   true,
		Messagge: message,
	}
}

func NewCreated(message string, data any) error {
	return &Response{
		Code: 201,
		Status: true,
		Messagge: message,
		Data: data,
	}
}

func NewBadRequestError(message string) error {
	return &ErrorResponse{
		Code: 400,
		Status: false,
		Messagge: message,
	}
}

func NewNotFoundError(message string) error {
	return &ErrorResponse{
		Code: 404,
		Status: false,
		Messagge: message,
	}
}

func NewValidationError(message string, data any) error {
	return &ErrorResponse{
		Code: 400,
		Status: false,
		Messagge: message,
		Err: data,
	}
}

func NewInternalServerError() error {
	return &ErrorResponse{
		Code: 500,
		Status: false,
		Messagge: "INTERNAL SERVER ERROR",
	}
}