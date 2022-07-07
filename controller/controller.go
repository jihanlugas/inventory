package controller

import (
	"context"
	"encoding/binary"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/constant"
	"github.com/jihanlugas/inventory/cryption"
	"github.com/jihanlugas/inventory/log"
	"github.com/jihanlugas/inventory/model"
	"github.com/jihanlugas/inventory/response"
	"github.com/jihanlugas/inventory/validator"
	"github.com/labstack/echo/v4"
	"net/http"
	"sync/atomic"
	"time"
)

type UserLogin struct {
	UserID      int64
	UserType    string
	PassVersion int
	PropertyID  int64
}

var CacheUserAuth map[int64]atomic.Value
var Validate *validator.CustomValidator

type UserAuthToken struct {
	PassVersion int
	LastLogin   int64
}

func init() {
	Validate = validator.NewValidator()
	CacheUserAuth = make(map[int64]atomic.Value)
}

func getLoginToken(userID int64, userType string, lastLoginDate int64, passVersion int, expiredAt time.Time, propertyID int64) ([]byte, error) {
	userType = fmt.Sprintf("%-20.20s", userType)
	expiredUnix := expiredAt.Unix()

	tokenPayload := make([]byte, constant.TokenPayloadLen)
	binary.BigEndian.PutUint64(tokenPayload, uint64(expiredUnix)) // Expired date
	copy(tokenPayload[8:32], userType)
	binary.BigEndian.PutUint64(tokenPayload[32:40], uint64(lastLoginDate))
	binary.BigEndian.PutUint64(tokenPayload[40:48], uint64(userID))
	binary.BigEndian.PutUint32(tokenPayload[48:52], uint32(passVersion))
	binary.BigEndian.PutUint64(tokenPayload[52:60], uint64(propertyID))

	return cryption.EncryptAES64(tokenPayload)
}

func getUserLoginInfo(c echo.Context) (UserLogin, error) {
	if u, ok := c.Get(constant.TokenUserContext).(UserLogin); ok {
		return u, nil
	} else {
		return UserLogin{}, response.ErrorForce("Akses tidak diterima", response.Payload{})
	}
}

func generateCookie(Name, Token string, ExpiredAt time.Time, MaxAge int) *http.Cookie {
	cookie := new(http.Cookie)
	cookie.Name = Name
	cookie.Value = Token
	cookie.Path = "/"
	cookie.Expires = ExpiredAt
	cookie.MaxAge = MaxAge
	cookie.SameSite = http.SameSiteNoneMode
	cookie.HttpOnly = false
	cookie.Secure = true

	return cookie
}

func genCacheUserAuth(userID int64, passVersion int, lastLogin int64) {
	if auth, ok := CacheUserAuth[userID]; ok {
		usrAuthVersion := auth.Load().(*UserAuthToken)
		usrAuthVersion.PassVersion = passVersion
		usrAuthVersion.LastLogin = lastLogin
		auth.Store(usrAuthVersion)
	} else {
		usrAuthVersion := &UserAuthToken{
			PassVersion: passVersion,
			LastLogin:   lastLogin,
		}
		var authUsr atomic.Value
		authUsr.Store(usrAuthVersion)
		CacheUserAuth[userID] = authUsr
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

func getPhotoUrl(ctx context.Context, conn *pgxpool.Conn, photoID int64) string {
	var err error
	var photo model.PublicPhoto
	if photoID == 0 {
		return ""
	}

	photo.PhotoID = photoID
	err = photo.GetById(ctx, conn)
	if err != nil {
		return ""
	}

	return config.PhotoAccessUrl + "/" + photo.PhotoPath
}
