package response

import (
	"encoding/json"
	"github.com/labstack/echo/v4"
	"net/http"
)

// SuccessResponse type for Success Response
type Response struct {
	Code      int         `json:"code"`
	IsSuccess bool        `json:"success"`
	Message   string      `json:"message"`
	Payload   interface{} `json:"payload" swaggertype:"object"`
}

type Payload map[string]interface{}

func (e *Response) Error() string {
	return e.Message
}

func StatusOk(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusOK,
		IsSuccess: true,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusCreated(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusCreated,
		IsSuccess: true,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusAccepted(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusAccepted,
		IsSuccess: true,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusBadRequest(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusBadRequest,
		IsSuccess: false,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusUnauthorized(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusUnauthorized,
		IsSuccess: false,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusForbidden(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusForbidden,
		IsSuccess: false,
		Message:   msg,
		Payload:   payload,
	}
}

func StatusNotFound(msg string, payload interface{}) *Response {
	return &Response{
		Code:      http.StatusNotFound,
		IsSuccess: false,
		Message:   msg,
		Payload:   payload,
	}
}

// ErrorForce generate ErrorResponse error type with additional forceLogout:true in payload
func ErrorForce(msg string, payload Payload) *Response {
	payload["forceLogout"] = true
	return &Response{
		Code:      http.StatusUnauthorized,
		IsSuccess: false,
		Message:   msg,
		Payload:   payload,
	}
}

func (r *Response) SendJSON(c echo.Context) error {
	return sendJSON(c, r, r.Code)
}

func sendJSON(c echo.Context, i interface{}, httpStatus int) error {
	if js, err := json.Marshal(i); err != nil {
		panic(err)
	} else {
		return c.Blob(httpStatus, echo.MIMEApplicationJSONCharsetUTF8, js)
	}
}

func ValidationError(err error) *Payload {
	return &Payload{
		"listError": getListError(err),
	}
}
