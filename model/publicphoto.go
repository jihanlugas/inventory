//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"time"
)

type PublicPhoto struct {
	PhotoID     int64      `db:"photo_id,pk" json:"photoId" form:"photoId" validate:"required"`
	ClientName  string     `db:"client_name,use_zero" json:"clientName" form:"clientName" validate:"required,lte=200"`
	ServerName  string     `db:"server_name,use_zero" json:"serverName" form:"serverName" validate:"required,lte=200"`
	Ext         string     `db:"ext,use_zero" json:"ext" form:"ext" validate:"required,lte=5"`
	PhotoPath   string     `db:"photo_path,use_zero" json:"photoPath" form:"photoPath" validate:"required,lte=200"`
	PhotoSize   int64      `db:"photo_size,use_zero" json:"photoSize" form:"photoSize" validate:"required"`
	PhotoWidth  int64      `db:"photo_width,use_zero" json:"photoWidth" form:"photoWidth" validate:"required"`
	PhotoHeight int64      `db:"photo_height,use_zero" json:"photoHeight" form:"photoHeight" validate:"required"`
	CreateBy    int64      `db:"create_by,use_zero" json:"createBy" form:"createBy" validate:"required"`
	CreateDt    *time.Time `db:"create_dt,use_zero" json:"createDt" form:"createDt" validate:"required"`
}
