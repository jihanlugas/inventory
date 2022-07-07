//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"context"
	"errors"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/config"
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/utils"
	"strconv"
)

func GetPhotoincQuery() *db.QueryComposer {
	return db.Query(`SELECT photoinc_id, ref_table, folder_inc, folder_name, running FROM public.photoinc`)
}

func (p *PublicPhotoinc) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	err = tx.QueryRow(ctx, `INSERT INTO public.photoinc
		(ref_table, folder_inc, folder_name, running)
		VALUES ($1, $2, $3, $4)
		RETURNING photoinc_id;`,
		p.RefTable,
		p.FolderInc,
		p.FolderName,
		p.Running,
	).Scan(&p.PhotoincID)

	return err
}

func (p *PublicPhotoinc) Update(ctx context.Context, tx pgx.Tx) error {
	var err error

	_, err = tx.Exec(ctx, `UPDATE public.photoinc SET ref_table = $1
		, folder_inc = $2
		, folder_name = $3
		, running = $4
		WHERE photoinc_id = $5`,
		p.RefTable,
		p.FolderInc,
		p.FolderName,
		p.Running,
		p.PhotoincID,
	)
	return err
}

func (p *PublicPhotoinc) GetTouse(conn *pgxpool.Conn, ctx context.Context, tx pgx.Tx, refTable string) error {
	var err error

	sql := GetPhotoincQuery().
		Where().
		StringEq("ref_table", refTable).
		Order("folder_inc", "DESC").
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			p.RefTable = refTable
			p.FolderInc = 1
			p.FolderName = config.PhotoDirectory + "/" + p.RefTable + "/" + strconv.FormatInt(p.FolderInc, 10)
			p.Running = 0
			err = p.Insert(ctx, tx)
			if err != nil {
				return err
			}

			err = utils.CreateFolder(p.FolderName, 0755)
			if err != nil {
				return err
			}
		} else {
			return err
		}
	} else {
		if p.Running >= config.PhotoincRunningLimit {
			p.RefTable = refTable
			p.FolderInc = p.FolderInc + 1
			p.FolderName = config.PhotoDirectory + "/" + p.RefTable + "/" + strconv.FormatInt(p.FolderInc, 10)
			p.Running = 0
			err = p.Insert(ctx, tx)
			if err != nil {
				return err
			}

			err = utils.CreateFolder(p.FolderName, 0755)
			if err != nil {
				return err
			}
		}
	}

	return err
}

func (p *PublicPhotoinc) AddRunning(ctx context.Context, tx pgx.Tx) error {
	var err error

	p.Running = p.Running + 1
	err = p.Update(ctx, tx)

	return err
}
