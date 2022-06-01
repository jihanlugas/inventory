//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/db"
	"time"
)

type PropertyRes struct {
	PropertyID   int64      `json:"propertyId"`
	PropertyName string     `json:"propertyName"`
	PhotoID      int64      `json:"photoId"`
	IsActive     bool       `json:"isActive"`
	CreateBy     int64      `json:"createBy"`
	CreateDt     *time.Time `json:"createDt"`
	UpdateBy     int64      `json:"updateBy"`
	UpdateDt     *time.Time `json:"updateDt"`
}

func GetPropertyQuery() *db.QueryComposer {
	return db.Query(`SELECT property_id, property_name, photo_id, is_active, create_by, create_dt, update_by, update_dt FROM public.property`)
}

func (p *PublicProperty) GetById(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetPropertyQuery().
		Where().
		Int64(`property_id`, "=", p.PropertyID).
		IsNull(`delete_dt`).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}

func (p *PublicProperty) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.CreateDt = &now
	p.UpdateDt = &now
	err = tx.QueryRow(ctx, `INSERT INTO public.property (property_name, photo_id, is_active, create_by, create_dt, update_by, update_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7)
		RETURNING property_id;`,
		p.PropertyName,
		p.PhotoID,
		p.IsActive,
		p.CreateBy,
		p.CreateDt,
		p.UpdateBy,
		p.UpdateDt,
	).Scan(&p.PropertyID)
	return err
}

func (p *PublicProperty) Res() PropertyRes {
	var res PropertyRes

	res.PropertyID = p.PropertyID
	res.PropertyName = p.PropertyName
	res.PhotoID = p.PhotoID
	res.IsActive = p.IsActive
	res.CreateBy = p.CreateBy
	res.CreateDt = p.CreateDt
	res.UpdateBy = p.UpdateBy
	res.UpdateDt = p.UpdateDt

	return res
}
