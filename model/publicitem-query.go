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

type ItemRes struct {
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
}

func GetItemQuery() *db.QueryComposer {
	return db.Query(`SELECT item_id, property_id, item_name, item_description, price, is_active, create_by, create_dt, update_by, update_dt FROM public.item`)
}

func GetItemWhere(ctx context.Context, conn *pgxpool.Conn, q *db.QueryBuilder) ([]PublicItem, error) {
	var err error
	var data []PublicItem

	err = pgxscan.Select(ctx, conn, &data, q.Build(), q.Params()...)
	if err != nil {
		return data, err
	}
	if len(data) == 0 {
		data = make([]PublicItem, 0)
	}

	return data, err
}

func (p *PublicItem) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.CreateDt = &now
	p.UpdateDt = &now
	err = tx.QueryRow(ctx, `INSERT INTO public.item
		(property_id, item_name, item_description, price, is_active, create_by, create_dt, update_by, update_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING item_id;`,
		p.PropertyID,
		p.ItemName,
		p.ItemDescription,
		p.Price,
		p.IsActive,
		p.CreateBy,
		p.CreateDt,
		p.UpdateBy,
		p.UpdateDt,
	).Scan(&p.ItemID)

	return err
}

func (p *PublicItem) Res() ItemRes {
	var res ItemRes

	res.ItemID = p.ItemID
	res.PropertyID = p.PropertyID
	res.ItemName = p.ItemName
	res.ItemDescription = p.ItemDescription
	res.Price = p.Price
	res.IsActive = p.IsActive
	res.CreateBy = p.CreateBy
	res.CreateDt = p.CreateDt
	res.UpdateBy = p.UpdateBy
	res.UpdateDt = p.UpdateDt

	return res
}
