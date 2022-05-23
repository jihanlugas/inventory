//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type PublicProperty struct {
	PropertyID   int64      `db:"property_id,pk" json:"propertyId" form:"propertyId" validate:"required"`
	PropertyName string     `db:"property_name,use_zero" json:"propertyName" form:"propertyName" validate:"required,lte=200"`
	PhotoID      int64      `db:"photo_id,use_zero" json:"photoId" form:"photoId" validate:"required"`
	IsActive     bool       `db:"is_active,use_zero" json:"isActive" form:"isActive" validate:"required"`
	CreateBy     int64      `db:"create_by,use_zero" json:"createBy" form:"createBy" validate:"required"`
	CreateDt     *time.Time `db:"create_dt,use_zero" json:"createDt" form:"createDt" validate:"required"`
	UpdateBy     int64      `db:"update_by,use_zero" json:"updateBy" form:"updateBy" validate:"required"`
	UpdateDt     *time.Time `db:"update_dt,use_zero" json:"updateDt" form:"updateDt" validate:"required"`
	DeleteBy     int64      `db:"delete_by" json:"deleteBy" form:"deleteBy"`
	DeleteDt     *time.Time `db:"delete_dt" json:"deleteDt" form:"deleteDt"`
}
