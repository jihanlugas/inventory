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

func GetItemvariantQuery() *db.QueryComposer {
	return db.Query(`SELECT itemvariant_id, item_id, itemvariant_name, itemvariant_description, price, is_active, photo_id, create_by, create_dt, update_by, update_dt FROM public.itemvariant`)
}

func (p *PublicItemvariant) GetById(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetItemvariantQuery().
		Where().
		Int64(`itemvariant_id`, "=", p.ItemvariantID).
		IsNull(`delete_dt`).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}

func GetItemvariantWhere(ctx context.Context, conn *pgxpool.Conn, q *db.QueryBuilder) ([]PublicItemvariant, error) {
	var err error
	var data []PublicItemvariant

	err = pgxscan.Select(ctx, conn, &data, q.Build(), q.Params()...)
	if err != nil {
		return data, err
	}
	if len(data) == 0 {
		data = make([]PublicItemvariant, 0)
	}

	return data, err
}

func (p *PublicItemvariant) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.CreateDt = &now
	p.UpdateDt = &now
	err = tx.QueryRow(ctx, `INSERT INTO public.itemvariant
		(item_id, itemvariant_name, itemvariant_description, price, photo_id, is_active, create_by, create_dt, update_by, update_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10)
		RETURNING itemvariant_id;`,
		p.ItemID,
		p.ItemvariantName,
		p.ItemvariantDescription,
		p.Price,
		p.PhotoID,
		p.IsActive,
		p.CreateBy,
		p.CreateDt,
		p.UpdateBy,
		p.UpdateDt,
	).Scan(&p.ItemvariantID)

	return err
}

func (p *PublicItemvariant) Update(ctx context.Context, tx pgx.Tx) error {

	var err error

	now := time.Now()
	p.UpdateDt = &now
	_, err = tx.Exec(ctx, `UPDATE public.itemvariant SET item_id = $1
		, itemvariant_name = $2
		, itemvariant_description = $3
		, price = $4
		, photo_id = $5
		, is_active = $6
		, update_by = $7
		, update_dt = $8
		WHERE itemvariant_id = $9`,
		p.ItemID,
		p.ItemvariantName,
		p.ItemvariantDescription,
		p.Price,
		p.PhotoID,
		p.IsActive,
		p.UpdateBy,
		p.UpdateDt,
		p.ItemvariantID,
	)
	return err
}
