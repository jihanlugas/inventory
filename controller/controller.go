package controller

import (
	"encoding/binary"
	"fmt"
	"github.com/jihanlugas/inventory/constant"
	"github.com/jihanlugas/inventory/cryption"
	"github.com/jihanlugas/inventory/log"
	"github.com/jihanlugas/inventory/response"
	"github.com/labstack/echo/v4"
	"time"
)

type UserLogin struct {
	UserID      int64
	UserType    string
	PassVersion int
}

func getLoginToken(userID int64, userType string, lastLoginDate int64, passVersion int, expiredAt time.Time) ([]byte, error) {
	userType = fmt.Sprintf("%-20.20s", userType)
	expiredUnix := expiredAt.Unix()

	tokenPayload := make([]byte, constant.TokenPayloadLen)
	binary.BigEndian.PutUint64(tokenPayload, uint64(expiredUnix)) // Expired date
	copy(tokenPayload[8:28], []byte(userType))
	binary.BigEndian.PutUint64(tokenPayload[28:], uint64(lastLoginDate))
	binary.BigEndian.PutUint64(tokenPayload[36:], uint64(userID))
	binary.BigEndian.PutUint32(tokenPayload[44:], uint32(passVersion))

	return cryption.EncryptAES64(tokenPayload)
}

func getUserLoginInfo(c echo.Context) (UserLogin, error) {
	if u, ok := c.Get(constant.TokenUserContext).(UserLogin); ok {
		return u, nil
	} else {
		return UserLogin{}, response.ErrorForce("Akses tidak diterima", response.Payload{})
	}
}

// Ping godoc
// @Summary      Ping
// // @Description  Ping server
// @Tags         Ping
// @Accept       json
// @Produce      json
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router       / [get]
func Ping(c echo.Context) error {
	return response.StatusOk("ようこそ、美しい世界へ", response.Payload{}).SendJSON(c)
}

func errorInternal(c echo.Context, err error) {
	log.System.Error().Err(err).Str("Host", c.Request().Host).Str("Path", c.Path()).Send()
	panic(err)
}
