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

func GetItemQuery() *db.QueryComposer {
	return db.Query(`SELECT item_id, property_id, item_name, item_description, is_active, photo_id, create_by, create_dt, update_by, update_dt FROM public.item`)
}

func (p *PublicItem) GetById(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetItemQuery().
		Where().
		Int64(`item_id`, "=", p.ItemID).
		IsNull(`delete_dt`).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
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
		(property_id, item_name, item_description, photo_id, is_active, create_by, create_dt, update_by, update_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING item_id;`,
		p.PropertyID,
		p.ItemName,
		p.ItemDescription,
		p.PhotoID,
		p.IsActive,
		p.CreateBy,
		p.CreateDt,
		p.UpdateBy,
		p.UpdateDt,
	).Scan(&p.ItemID)

	return err
}

func (p *PublicItem) Update(ctx context.Context, tx pgx.Tx) error {

	var err error

	now := time.Now()
	p.UpdateDt = &now
	_, err = tx.Exec(ctx, `UPDATE public.item SET property_id = $1
		, item_name = $2
		, item_description = $3
		, photo_id = $4
		, is_active = $5
		, update_by = $6
		, update_dt = $7
		WHERE item_id = $8`,
		p.PropertyID,
		p.ItemName,
		p.ItemDescription,
		p.PhotoID,
		p.IsActive,
		p.UpdateBy,
		p.UpdateDt,
		p.ItemID,
	)
	return err
}

//func (p *PublicItem) Res(ctx context.Context, conn *pgxpool.Conn) ItemRes {
//	var res ItemRes
//
//	res.ItemID = p.ItemID
//	res.PropertyID = p.PropertyID
//	res.ItemName = p.ItemName
//	res.ItemDescription = p.ItemDescription
//	res.Price = p.Price
//	res.PhotoUrl = controller.GetPhotoUrl(ctx, conn, p.PhotoID)
//	res.IsActive = p.IsActive
//	res.CreateBy = p.CreateBy
//	res.CreateDt = p.CreateDt
//	res.UpdateBy = p.UpdateBy
//	res.UpdateDt = p.UpdateDt
//
//	return res
//}
