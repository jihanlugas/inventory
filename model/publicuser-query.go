package model

import (
	"context"
	"github.com/georgysavva/scany/pgxscan"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/jihanlugas/inventory/db"
	"strings"
	"time"
)

type UserRes struct {
	UserID   int64  `json:"userId"`
	Fullname string `json:"fullname"`
	NoHp     string `json:"noHp"`
	Email    string `json:"email"`
	UserType string `json:"userType"`
	PhotoID  int64  `json:"photoId"`
	IsActive bool   `json:"isActive"`
	PhotoUrl string `json:"photoUrl"`
}

func GetUserQuery() *db.QueryComposer {
	return db.Query(`SELECT user_id, fullname, email, no_hp, user_type, username, passwd, photo_id, is_active, last_login_dt, pass_version, create_by, create_dt, update_by, update_dt FROM public.user`)
}

func (p *PublicUser) GetById(ctx context.Context, conn *pgxpool.Conn) error {
	var err error

	sql := GetUserQuery().
		Where().
		Int64(`user_id`, "=", p.UserID).
		IsNull(`delete_dt`).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}

func (p *PublicUser) GetByUsername(ctx context.Context, conn *pgxpool.Conn) error {
	var err error
	p.Username = strings.ToLower(p.Username)

	sql := GetUserQuery().
		Where().
		StringEq(`username`, p.Username).
		IsNull(`delete_dt`).
		OffsetLimit(0, 1)
	err = pgxscan.Get(ctx, conn, p, sql.Build(), sql.Params()...)

	return err
}

func (p *PublicUser) Insert(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.Username = strings.ToLower(p.Username)
	p.CreateDt = &now
	p.UpdateDt = &now
	err = tx.QueryRow(ctx, `INSERT INTO public.user (fullname, email, no_hp, user_type, username, passwd, photo_id, is_active, last_login_dt, pass_version, create_by, create_dt, update_by, update_dt)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14)
		RETURNING user_id;`,
		p.Fullname,
		p.Email,
		p.NoHp,
		p.UserType,
		p.Username,
		p.Passwd,
		p.PhotoID,
		p.IsActive,
		p.LastLoginDt,
		p.PassVersion,
		p.CreateBy,
		p.CreateDt,
		p.UpdateBy,
		p.UpdateDt,
	).Scan(&p.UserID)
	return err
}

// passwd not updated
func (p *PublicUser) Update(ctx context.Context, tx pgx.Tx) error {
	var err error

	now := time.Now()
	p.Username = strings.ToLower(p.Username)
	p.UpdateDt = &now
	_, err = tx.Exec(ctx, `UPDATE public.user SET fullname = $1
		, email = $2
		, no_hp = $3
		, user_type = $4
		, username = $5
		, photo_id = $6
		, is_active = $7
		, last_login_dt = $8
		, pass_version = $9
		, update_by = $10
		, update_dt = $11
		WHERE user_id = $12`,
		p.Fullname,
		p.Email,
		p.NoHp,
		p.UserType,
		p.Username,
		p.PhotoID,
		p.IsActive,
		p.LastLoginDt,
		p.PassVersion,
		p.UpdateBy,
		p.UpdateDt,
		p.UserID,
	)
	return err
}

func (p *PublicUser) Res() UserRes {
	var res UserRes

	res.UserID = p.UserID
	res.Fullname = p.Fullname
	res.NoHp = p.NoHp
	res.Email = p.Email
	res.UserType = p.UserType
	res.PhotoID = p.PhotoID
	res.IsActive = p.IsActive

	return res
}
