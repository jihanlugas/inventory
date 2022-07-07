//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type PublicItemvariant struct {
	ItemvariantID          int64      `db:"itemvariant_id,pk" json:"itemvariantId" form:"itemvariantId" validate:"required"`
	ItemID                 int64      `db:"item_id,use_zero" json:"itemId" form:"itemId" validate:"required"`
	ItemvariantName        string     `db:"itemvariant_name,use_zero" json:"itemvariantName" form:"itemvariantName" validate:"required,lte=200"`
	ItemvariantDescription string     `db:"itemvariant_description,use_zero" json:"itemvariantDescription" form:"itemvariantDescription" validate:"required,lte=500"`
	Price                  int64      `db:"price,use_zero" json:"price" form:"price" validate:"required"`
	PhotoID                int64      `db:"photo_id,use_zero" json:"photoId" form:"photoId" validate:"required"`
	IsActive               bool       `db:"is_active,use_zero" json:"isActive" form:"isActive" validate:"required"`
	CreateBy               int64      `db:"create_by,use_zero" json:"createBy" form:"createBy" validate:"required"`
	CreateDt               *time.Time `db:"create_dt,use_zero" json:"createDt" form:"createDt" validate:"required"`
	UpdateBy               int64      `db:"update_by,use_zero" json:"updateBy" form:"updateBy" validate:"required"`
	UpdateDt               *time.Time `db:"update_dt,use_zero" json:"updateDt" form:"updateDt" validate:"required"`
	DeleteBy               int64      `db:"delete_by" json:"deleteBy" form:"deleteBy"`
	DeleteDt               *time.Time `db:"delete_dt" json:"deleteDt" form:"deleteDt"`
}
