package controller

import (
	"errors"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/constant"
	"github.com/jihanlugas/inventory/cryption"
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/model"
	"github.com/jihanlugas/inventory/response"
	"github.com/jihanlugas/inventory/utils"
	"github.com/labstack/echo/v4"
	"time"
)

type User struct{}

func UserComposer() User {
	return User{}
}

type signupReq struct {
	Fullname      string `json:"fullname" form:"fullname" validate:"required,lte=80"`
	Email         string `json:"email" form:"email" validate:"required,lte=200,email,notexists=email email"`
	NoHp          string `json:"noHp" form:"noHp" validate:"required,lte=20,notexists=no_hp noHp"`
	Username      string `json:"username" form:"username" validate:"required,lte=20,lowercase,notexists=username username"`
	Passwd        string `json:"passwd" form:"passwd" validate:"required,lte=200"`
	ConfirmPasswd string `json:"confirmPasswd" form:"confirmPasswd" validate:"required,lte=200,eqfield=Passwd"`
	PropertyName  string `json:"propertyName" form:"propertyName" validate:"required,lte=200"`
}

type signinReq struct {
	Username string `db:"username,use_zero" json:"username" form:"username" query:"username" validate:"required,lte=20"`
	Passwd   string `db:"passwd,use_zero" json:"passwd" form:"passwd" query:"passwd" validate:"required,lte=200"`
}

type meRes struct {
	UserID       int64  `json:"userId"`
	Fullname     string `json:"fullname"`
	Email        string `json:"email"`
	Username     string `json:"username"`
	PropertyID   int64  `json:"propertyId"`
	PropertyName string `json:"propertyName"`
}

// SignUp Sign Up User
// @Summary Sign up a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body signupReq true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-up [post]
func (h User) SignUp(c echo.Context) error {
	var err error
	req := new(signupReq)
	//var photouser model.PublicPhoto

	if err = c.Bind(req); err != nil {
		return err
	}

	//req.PhotoUser, err = c.FormFile("photoUser")
	//if err != nil && !errors.Is(err, http.ErrMissingFile) {
	//	errorInternal(c, err)
	//}
	//req.PhotoUserChk = req.PhotoUser != nil

	if err = c.Validate(req); err != nil {
		fmt.Println("validation failed")
		return response.StatusBadRequest("validation failed", response.ValidationError(err)).SendJSON(c)
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	newPasswd, err := cryption.HashPassword(req.Passwd)
	if err != nil {
		fmt.Println("hash password failed")
		return response.StatusBadRequest("hash password failed", response.Payload{}).SendJSON(c)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		errorInternal(c, err)
	}
	defer db.DeferHandleTransaction(ctx, tx)

	now := time.Now()

	var user model.PublicUser
	user.Fullname = req.Fullname
	user.Email = req.Email
	user.UserType = constant.UserTypeAdmin
	user.Username = req.Username
	user.NoHp = utils.FormatPhoneTo62(req.NoHp)
	user.Passwd = newPasswd
	user.PhotoID = 0
	user.LastLoginDt = &now
	user.PassVersion = 1
	user.IsActive = true
	user.CreateBy = 0
	user.UpdateBy = 0
	err = user.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	var property model.PublicProperty
	property.PropertyName = req.PropertyName
	property.PhotoID = 0
	property.IsActive = true
	err = property.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	var userproperty model.PublicUserproperty
	userproperty.UserID = user.UserID
	userproperty.PropertyID = property.PropertyID
	userproperty.IsDefault = true
	err = userproperty.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	//if req.PhotoUser != nil {
	//	err = photouser.Upload(ctx, tx, req.PhotoUser, constant.FOTO_USER, user.UserID)
	//	if err != nil {
	//		errorInternal(c, err)
	//	}
	//}

	//user.PhotoID = photouser.PhotoID
	//err = user.Update(ctx, tx)
	//if err != nil {
	//	errorInternal(c, err)
	//}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		errorInternal(c, err)
	}

	return response.StatusCreated("Success Signup", response.Payload{
		"user":         user.Res(),
		"property":     property.Res(),
		"userproperty": userproperty,
	}).SendJSON(c)
}

// SignIn Sign In user
// @Summary Sign in a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body signinReq true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-in [post]
func (h User) SignIn(c echo.Context) error {
	var err error
	var ok bool

	req := new(signinReq)
	if err = c.Bind(req); err != nil {
		return err
	}

	if err = c.Validate(req); err != nil {
		return response.StatusBadRequest("error validation", response.ValidationError(err)).SendJSON(c)
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	var user model.PublicUser
	user.Username = req.Username
	err = user.GetByUsername(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusBadRequest("invalid username or password", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	if !user.IsActive {
		return response.StatusBadRequest("user not active", response.Payload{}).SendJSON(c)
	}

	if ok = cryption.CheckPasswordHash(req.Passwd, user.Passwd); !ok {
		return response.StatusBadRequest("invalid username or password", response.Payload{}).SendJSON(c)
	}

	var userproperty model.PublicUserproperty
	userproperty.UserID = user.UserID
	err = userproperty.GetDefaultProperty(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("user property not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	var property model.PublicProperty
	property.PropertyID = userproperty.PropertyID
	err = property.GetById(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("property not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	tx, err := conn.Begin(ctx)
	if err != nil {
		errorInternal(c, err)
	}
	defer db.DeferHandleTransaction(ctx, tx)

	now := time.Now()
	user.LastLoginDt = &now
	err = user.Update(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		errorInternal(c, err)
	}

	lastLoginNano := user.LastLoginDt.UnixNano()
	expiredAt := time.Now().Add(time.Hour * 12)
	token, err := getLoginToken(user.UserID, user.UserType, lastLoginNano, user.PassVersion, expiredAt, userproperty.PropertyID)
	if err != nil {
		return response.StatusBadRequest("generate token failed", response.Payload{}).SendJSON(c)
	}
	maxAge := expiredAt.Sub(now).Seconds()

	cookie := generateCookie(config.CookieAuthName, string(token), expiredAt, int(maxAge))
	c.SetCookie(cookie)

	genCacheUserAuth(user.UserID, user.PassVersion, lastLoginNano)

	res := meRes{
		UserID:       user.UserID,
		Fullname:     user.Fullname,
		Email:        user.Email,
		Username:     user.Username,
		PropertyID:   property.PropertyID,
		PropertyName: property.PropertyName,
	}

	return response.StatusOk("success", res).SendJSON(c)
}

// SignOut godoc
// @Summary Sign out a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-out [get]
func (h User) SignOut(c echo.Context) error {
	return response.StatusOk("Unhandled Function", response.Payload{}).SendJSON(c)
}

// Me godoc
// @Tags User
// @Summary To do get current active user
// @Accept json
// @Produce json
// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /user/me [get]
func (h User) Me(c echo.Context) error {
	var err error
	loginUser, err := getUserLoginInfo(c)
	if err != nil {
		return err
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	var user model.PublicUser
	user.UserID = loginUser.UserID
	err = user.GetById(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.ErrorForce("data not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	var userproperty model.PublicUserproperty
	userproperty.UserID = user.UserID
	err = userproperty.GetDefaultProperty(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("user property not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	var property model.PublicProperty
	property.PropertyID = userproperty.PropertyID
	err = property.GetById(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("property not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	res := meRes{
		UserID:       user.UserID,
		Fullname:     user.Fullname,
		Email:        user.Email,
		Username:     user.Username,
		PropertyID:   property.PropertyID,
		PropertyName: property.PropertyName,
	}

	return response.StatusOk("Success", res).SendJSON(c)
}
