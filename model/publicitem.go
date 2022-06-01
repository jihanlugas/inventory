//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type PublicItem struct {
	ItemID          int64      `db:"item_id,pk" json:"itemId" form:"itemId" validate:"required"`
	PropertyID      int64      `db:"property_id,use_zero" json:"propertyId" form:"propertyId" validate:"required"`
	ItemName        string     `db:"item_name,use_zero" json:"itemName" form:"itemName" validate:"required,lte=200"`
	ItemDescription string     `db:"item_description,use_zero" json:"itemDescription" form:"itemDescription" validate:"required,lte=200"`
	Price           int64      `db:"price,use_zero" json:"price" form:"price" validate:"required"`
	IsActive        bool       `db:"is_active,use_zero" json:"isActive" form:"isActive" validate:"required"`
	CreateBy        int64      `db:"create_by,use_zero" json:"createBy" form:"createBy" validate:"required"`
	CreateDt        *time.Time `db:"create_dt,use_zero" json:"createDt" form:"createDt" validate:"required"`
	UpdateBy        int64      `db:"update_by,use_zero" json:"updateBy" form:"updateBy" validate:"required"`
	UpdateDt        *time.Time `db:"update_dt,use_zero" json:"updateDt" form:"updateDt" validate:"required"`
	DeleteBy        int64      `db:"delete_by" json:"deleteBy" form:"deleteBy"`
	DeleteDt        *time.Time `db:"delete_dt" json:"deleteDt" form:"deleteDt"`
}
