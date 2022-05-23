package controller

import (
	"github.com/jihanlugas/inventory/response"
	"github.com/labstack/echo/v4"
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
}

type signinReq struct {
	Username string `db:"username,use_zero" json:"username" form:"username" query:"username" validate:"required,lte=20"`
	Passwd   string `db:"passwd,use_zero" json:"passwd" form:"passwd" query:"passwd" validate:"required,lte=200"`
}

// SignUp godoc
// @Summary Sign up a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body signupReq true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-up [post]
func (h User) SignUp(c echo.Context) error {
	return response.StatusOk("Unhandled Function", response.Payload{}).SendJSON(c)
}

// SignIn godoc
// @Summary Sign in a user
// @Tags Authentication
// @Accept json
// @Produce json
// @Param req body signinReq true "json req body"
// @Success      200  {object}	response.Response
// @Failure      500  {object}  response.Response
// @Router /sign-in [post]
func (h User) SignIn(c echo.Context) error {
	//var err error
	//var ok bool
	//
	//req := new(signinReq)
	//if err = c.Bind(req); err != nil {
	//	return err
	//}
	//
	//if err = c.Validate(req); err != nil {
	//	return response.StatusBadRequest("error validation", response.ValidationError(err)).SendJSON(c)
	//}
	//
	//conn, ctx, closeConn := db.GetConnection()
	//defer closeConn()
	//
	//var user model.PublicUser
	//user.Username = req.Username
	//err = user.GetByUsername(ctx, conn)
	//if err != nil {
	//	if errors.Is(err, pgx.ErrNoRows) {
	//		return response.StatusBadRequest("invalid username or password", response.Payload{}).SendJSON(c)
	//	}
	//	errorInternal(c, err)
	//}
	//
	//if !user.IsActive {
	//	return response.StatusBadRequest("User not active", response.Payload{}).SendJSON(c)
	//}
	//
	//if ok = cryption.CheckPasswordHash(req.Passwd, user.Passwd); !ok {
	//	return response.StatusBadRequest("invalid username or password", response.Payload{}).SendJSON(c)
	//}
	//
	//tx, err := conn.Begin(ctx)
	//if err != nil {
	//	errorInternal(c, err)
	//}
	//defer db.DeferHandleTransaction(ctx, tx)
	//
	//now := time.Now()
	//user.LastLoginDt = &now
	//err = user.Update(ctx, tx)
	//if err != nil {
	//	errorInternal(c, err)
	//}
	//
	//if err = tx.Commit(ctx); err != nil {
	//	_ = tx.Rollback(ctx)
	//	errorInternal(c, err)
	//}
	//
	//lastLoginNano := user.LastLoginDt.UnixNano()
	//expiredAt := time.Now().Add(time.Hour * 12)
	//token, err := getLoginToken(int64(user.UserID), user.UserType, lastLoginNano, int(user.PassVersion), expiredAt)
	//if err != nil {
	//	return response.StatusBadRequest("generate token failed", response.Payload{}).SendJSON(c)
	//}
	//maxAge := expiredAt.Sub(now).Seconds()
	//
	//cookie := generateCookie(config.CookieAuthName, string(token), expiredAt, int(maxAge))
	//c.SetCookie(cookie)
	//
	//genCacheUserAuth(int64(user.UserID), int(user.PassVersion), lastLoginNano)
	//
	//res := meRes{
	//	UserID:   user.UserID,
	//	Fullname: user.Fullname,
	//	Email:    user.Email,
	//	Username: user.Username,
	//}

	return response.StatusOk("Unhandled Function", response.Payload{}).SendJSON(c)
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
