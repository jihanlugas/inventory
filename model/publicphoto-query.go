//nolint
//lint:file-ignore U1000 ignore unused code, it's generated
package model

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/utils"
	"mime/multipart"
	"path/filepath"
	"strconv"
	"time"
)

func GetPhotoQuery() *db.QueryComposer {
	return db.Query(`SELECT photo_id, client_name, server_name, ext, photo_path, photo_size, photo_width, photo_height, create_by, create_dt FROM public.photo`)
}

func (p *PublicPhoto) GetById(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetPhotoQuery().
		Where().
		Int64(`photo_id`, "=", p.PhotoID).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}

func (p *PublicPhoto) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.CreateDt = &now
	err = tx.QueryRow(ctx, `INSERT INTO public.photo
		(client_name, server_name, ext, photo_path, photo_size, photo_width, photo_height, create_by, create_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING photo_id;`,
		p.ClientName,
		p.ServerName,
		p.Ext,
		p.PhotoPath,
		p.PhotoSize,
		p.PhotoWidth,
		p.PhotoHeight,
		p.CreateBy,
		p.CreateDt,
	).Scan(&p.PhotoID)

	if err != nil {
		return err
	}

	p.ServerName = strconv.FormatInt(p.PhotoID, 10) + p.Ext
	p.PhotoPath = p.PhotoPath + "/" + p.ServerName
	_, err = tx.Exec(ctx, `UPDATE public.photo SET server_name = $1
		, photo_path = $2
		WHERE photo_id = $3`,
		p.ServerName,
		p.PhotoPath,
		p.PhotoID,
	)

	return err
}

func (p *PublicPhoto) Upload(conn *pgxpool.Conn, ctx context.Context, tx pgx.Tx, file *multipart.FileHeader, refTable string) error {
	var err error
	var publicphotoinc PublicPhotoinc

	err = publicphotoinc.GetTouse(conn, ctx, tx, refTable)
	if err != nil {
		return err
	}

	err = publicphotoinc.AddRunning(ctx, tx)
	if err != nil {
		return err
	}

	now := time.Now()
	p.ClientName = file.Filename
	p.Ext = filepath.Ext(file.Filename)
	p.PhotoPath = publicphotoinc.FolderName
	p.PhotoSize = file.Size
	p.PhotoWidth = 0
	p.PhotoHeight = 0
	p.CreateDt = &now
	err = p.Insert(ctx, tx)
	if err != nil {
		return err
	}

	err = utils.UploadImage(p.PhotoPath, file)
	if err != nil {
		return err
	}

	return nil
}

func (p *PublicPhoto) Delete(conn *pgxpool.Conn, ctx context.Context, tx pgx.Tx) error {
	var err error

	err = p.GetById(ctx, conn)
	if err != nil {
		return err
	}

	err = utils.RemoveImage(p.PhotoPath)
	if err != nil {
		return err
	}

	_, err = tx.Exec(ctx, "DELETE FROM public.photo WHERE photo_id = $1", p.PhotoID)
	if err != nil {
		return err
	}

	return err
}
