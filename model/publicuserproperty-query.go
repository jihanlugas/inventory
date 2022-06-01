//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/db"
)

func GetUserpropertyQuery() *db.QueryComposer {
	return db.Query(`SELECT user_id, property_id, is_default FROM public.userproperty`)
}

func (p *PublicUserproperty) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	err = tx.QueryRow(ctx, `INSERT INTO public.userproperty
		(user_id, property_id, is_default)
		VALUES ($1, $2, $3)
		RETURNING userproperty_id;`,
		p.UserID,
		p.PropertyID,
		p.IsDefault,
	).Scan(&p.UserpropertyID)

	return err
}

func (p *PublicUserproperty) GetDefaultProperty(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetUserpropertyQuery().
		Where().
		Int64(`user_id`, "=", p.UserID).
		Bool("is_default", true).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}
