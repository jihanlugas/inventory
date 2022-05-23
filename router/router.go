package router

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"fmt"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/constant"
	"github.com/jihanlugas/inventory/controller"
	"github.com/jihanlugas/inventory/cryption"
	"github.com/jihanlugas/inventory/response"
	"github.com/labstack/echo/v4"
	"net/http"
	"time"

	_ "github.com/jihanlugas/inventory/swg"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func Init() *echo.Echo {

	web := websiteRouter()
	checkToken := checkTokenMiddleware()
	userController := controller.UserComposer()

	web.GET("/swg/*", echoSwagger.WrapHandler)
	web.GET("/", controller.Ping)
	web.POST("/sign-up", userController.SignUp)
	web.POST("/sign-in", userController.SignIn)
	web.GET("/sign-out", userController.SignOut, checkToken)

	return web
}

func checkTokenMiddleware() echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cookie, err := c.Cookie(config.CookieAuthName)
			if err != nil {
				return response.ErrorForce("Akses ditolak!", response.Payload{}).SendJSON(c)
			}

			token := cookie.Value
			tokenPayload, err := cryption.DecryptAES64([]byte(token))
			if err != nil {
				return response.ErrorForce("Akses telah kadaluarsa", response.Payload{}).SendJSON(c)
			}

			if len(tokenPayload) == constant.TokenPayloadLen {
				expiredUnix := binary.BigEndian.Uint64(tokenPayload)
				expiredAt := time.Unix(int64(expiredUnix), 0)
				now := time.Now()
				if now.After(expiredAt) {
					return response.ErrorForce("Akses telah kadaluarsa!", response.Payload{}).SendJSON(c)
				} else {
					usrLogin := controller.UserLogin{
						UserType:    string(bytes.TrimSpace(tokenPayload[8:28])),
						UserID:      int64(binary.BigEndian.Uint64(tokenPayload[36:])),
						PassVersion: int(binary.LittleEndian.Uint16(tokenPayload[44:])),
					}
					c.Set(constant.TokenUserContext, usrLogin)
					return next(c)
				}
			} else {
				return response.ErrorForce("Akses telah kadaluarsa!", response.Payload{}).SendJSON(c)
			}
		}
	}
}

func httpErrorHandler(err error, c echo.Context) {
	var errorResponse *response.Response
	code := http.StatusInternalServerError
	switch e := err.(type) {
	case *echo.HTTPError:
		// Handle pada saat URL yang di request tidak ada. atau ada kesalahan server.
		code = e.Code
		errorResponse = &response.Response{
			IsSuccess: false,
			Message:   fmt.Sprintf("%v", e.Message),
			Payload:   map[string]interface{}{},
			Code:      code,
		}
	case *response.Response:
		errorResponse = e
	default:
		// Handle error dari panic
		code = http.StatusInternalServerError
		if config.Environment != config.Production {
			errorResponse = &response.Response{
				IsSuccess: false,
				Message:   err.Error(),
				Payload:   map[string]interface{}{},
				Code:      http.StatusInternalServerError,
			}
		} else {
			errorResponse = &response.Response{
				IsSuccess: false,
				Message:   "Internal server error",
				Payload:   map[string]interface{}{},
				Code:      http.StatusInternalServerError,
			}
		}
	}

	js, err := json.Marshal(errorResponse)
	if err == nil {
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, js)
	} else {
		b := []byte("{error: true, message: \"unresolved error\"}")
		_ = c.Blob(code, echo.MIMEApplicationJSONCharsetUTF8, b)
	}
}
