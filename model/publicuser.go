//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type PublicUser struct {
	UserID      int64      `db:"user_id,pk" json:"userId" form:"userId" validate:"required"`
	Fullname    string     `db:"fullname,use_zero" json:"fullname" form:"fullname" validate:"required,lte=80"`
	NoHp        string     `db:"no_hp,use_zero" json:"noHp" form:"noHp" validate:"required,lte=20"`
	Email       string     `db:"email,use_zero" json:"email" form:"email" validate:"required,lte=200"`
	UserType    string     `db:"user_type,use_zero" json:"userType" form:"userType" validate:"required,oneof='ADMIN' 'USER'"`
	Username    string     `db:"username,use_zero" json:"username" form:"username" validate:"required,lte=20"`
	Passwd      string     `db:"passwd,use_zero" json:"passwd" form:"passwd" validate:"required,lte=200"`
	PhotoID     int64      `db:"photo_id,use_zero" json:"photoId" form:"photoId" validate:"required"`
	IsActive    bool       `db:"is_active,use_zero" json:"isActive" form:"isActive" validate:"required"`
	LastLoginDt *time.Time `db:"last_login_dt,use_zero" json:"lastLoginDt" form:"lastLoginDt" validate:"required"`
	PassVersion int        `db:"pass_version,use_zero" json:"passVersion" form:"passVersion" validate:"required"`
	CreateBy    int64      `db:"create_by,use_zero" json:"createBy" form:"createBy" validate:"required"`
	CreateDt    *time.Time `db:"create_dt,use_zero" json:"createDt" form:"createDt" validate:"required"`
	UpdateBy    int64      `db:"update_by,use_zero" json:"updateBy" form:"updateBy" validate:"required"`
	UpdateDt    *time.Time `db:"update_dt,use_zero" json:"updateDt" form:"updateDt" validate:"required"`
	DeleteBy    int64      `db:"delete_by" json:"deleteBy" form:"deleteBy"`
	DeleteDt    *time.Time `db:"delete_dt" json:"deleteDt" form:"deleteDt"`
}
