package controller

import (
	"errors"
	"github.com/jackc/pgx/v4"
	"github.com/jihanlugas/inventory/constant"
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/model"
	"github.com/jihanlugas/inventory/request"
	"github.com/jihanlugas/inventory/response"
	"github.com/labstack/echo/v4"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"
)

type Itemvariant struct{}

func ItemvariantComposer() Itemvariant {
	return Itemvariant{}
}

type pageItemvariantReq struct {
	request.Paging
	ItemID                 int64  `json:"itemId" form:"itemId" query:"itemId" validate:"required"`
	ItemvariantName        string `json:"itemvariantName" form:"itemvariantName" query:"itemvariantName"`
	ItemvariantDescription string `json:"itemvariantDescription" form:"itemvariantDescription" query:"itemvariantDescription"`
}

type createItemvariantReq struct {
	ItemID                 int64                 `json:"itemId" form:"itemId" validate:"required"`
	ItemvariantName        string                `json:"itemvariantName" form:"itemvariantName" validate:"required,lte=200"`
	ItemvariantDescription string                `json:"itemvariantDescription" form:"itemvariantDescription" validate:"lte=500"`
	Price                  int64                 `json:"price" form:"price" validate:"required"`
	IsActive               bool                  `json:"isActive" form:"isActive" validate:"required"`
	Photo                  *multipart.FileHeader `json:"-"`
	PhotoChk               bool                  `json:"photo" validate:"required,photo=Photo"`
}

type itemvariantRes struct {
	ItemvariantID          int64      `json:"itemvariantId"`
	ItemID                 int64      `json:"itemId"`
	ItemvariantName        string     `json:"itemvariantName"`
	ItemvariantDescription string     `json:"itemvariantDescription"`
	Price                  int64      `json:"price"`
	PhotoUrl               string     `json:"photoUrl"`
	IsActive               bool       `json:"isActive"`
	CreateBy               int64      `json:"createBy"`
	CreateDt               *time.Time `json:"createDt"`
	UpdateBy               int64      `json:"updateBy"`
	UpdateDt               *time.Time `json:"updateDt"`
}

// Create Create Itemvariant
// @Summary Create Itemvariant
// @Tags Itemvariant
// @Accept json
// @Produce json
// @Param itemvariantName formData string true "Itemvariant Name"
// @Param itemvariantDescription formData string true "Itemvariant Description"
// @Param price formData integer true "Price"
// @Param isActive formData boolean true "Active"
// @Param photo formData file true "Photo"
// @Success      200  {object}	response.Response{payload=itemvariantRes}
// @Failure      500  {object}  response.Response
// @Router /itemvariant [post]
func (h Itemvariant) Create(c echo.Context) error {
	var err error
	var data model.PublicItemvariant
	var publicphoto model.PublicPhoto

	loginUser, err := getUserLoginInfo(c)
	if err != nil {
		errorInternal(c, err)
	}

	req := new(createItemvariantReq)
	if err = c.Bind(req); err != nil {
		return err
	}

	req.Photo, err = c.FormFile("photo")
	if err != nil && !errors.Is(err, http.ErrMissingFile) {
		errorInternal(c, err)
	}
	req.PhotoChk = req.Photo != nil

	if err = c.Validate(req); err != nil {
		return response.StatusBadRequest("validation failed", response.ValidationError(err)).SendJSON(c)
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	tx, err := conn.Begin(ctx)
	if err != nil {
		errorInternal(c, err)
	}
	defer db.DeferHandleTransaction(ctx, tx)

	data.ItemID = req.ItemID
	data.ItemvariantName = req.ItemvariantName
	data.ItemvariantDescription = req.ItemvariantDescription
	data.Price = req.Price
	data.IsActive = req.IsActive
	data.CreateBy = loginUser.UserID
	data.UpdateBy = loginUser.UserID
	err = data.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	publicphoto.CreateBy = loginUser.UserID
	err = publicphoto.Upload(conn, ctx, tx, req.Photo, constant.PhotoRefTableItemvariant)
	if err != nil {
		errorInternal(c, err)
	}

	data.PhotoID = publicphoto.PhotoID
	data.UpdateBy = loginUser.UserID
	err = data.Update(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		errorInternal(c, err)
	}

	res := itemvariantRes{
		ItemvariantID:          data.ItemvariantID,
		ItemID:                 data.ItemID,
		ItemvariantName:        data.ItemvariantName,
		ItemvariantDescription: data.ItemvariantDescription,
		Price:                  data.Price,
		PhotoUrl:               getPhotoUrl(ctx, conn, data.PhotoID),
		IsActive:               data.IsActive,
		CreateBy:               data.CreateBy,
		CreateDt:               data.CreateDt,
		UpdateBy:               data.UpdateBy,
		UpdateDt:               data.UpdateDt,
	}

	return response.StatusCreated("success", res).SendJSON(c)
}

// GetById Get Itemvariant
// @Summary Get Itemvariant
// @Tags Itemvariant
// @Accept json
// @Produce json
// @Param kanji path number true "itemvariant_id" default(0)
// @Success      200  {object}	response.Response{payload=itemvariantRes}
// @Failure      500  {object}  response.Response
// @Router /itemvariant/{itemvariant_id} [get]
func (h Itemvariant) GetById(c echo.Context) error {
	var err error
	var data model.PublicItemvariant

	ID, err := strconv.ParseInt(c.Param("itemvariant_id"), 10, 64)
	if err != nil {
		return response.StatusNotFound("data not found", response.Payload{}).SendJSON(c)
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	data.ItemvariantID = ID
	err = data.GetById(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("data not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	var res itemvariantRes
	res.ItemID = data.ItemID
	res.ItemvariantID = data.ItemvariantID
	res.ItemvariantName = data.ItemvariantName
	res.ItemvariantDescription = data.ItemvariantDescription
	res.Price = data.Price
	res.PhotoUrl = getPhotoUrl(ctx, conn, data.PhotoID)
	res.IsActive = data.IsActive
	res.CreateBy = data.CreateBy
	res.CreateDt = data.CreateDt
	res.UpdateBy = data.UpdateBy
	res.UpdateDt = data.UpdateDt

	return response.StatusOk("success", res).SendJSON(c)
}

// Page Page Itemvariant
// @Summary Page Itemvariant
// @Tags Itemvariant
// @Accept json
// @Produce json
// @Param req body pageItemvariantReq true "payload"
// @Success      200  {object}	response.Response{payload=response.Pagination}
// @Failure      500  {object}  response.Response
// @Router /itemvariant/page [post]
func (h Itemvariant) Page(c echo.Context) error {
	var err error

	req := new(pageItemvariantReq)
	if err = c.Bind(req); err != nil {
		errorInternal(c, err)
	}

	if err = c.Validate(req); err != nil {
		return response.StatusBadRequest("validation error", response.ValidationError(err)).SendJSON(c)
	}

	err, cnt, list := getPageItemvariant(req)
	if err != nil {
		errorInternal(c, err)
	}

	return response.StatusOk("success", response.PayloadPagination(req, list, cnt)).SendJSON(c)
}

func getPageItemvariant(req *pageItemvariantReq) (error, int, []itemvariantRes) {
	var err error
	var cnt int
	var list []model.PublicItemvariant
	var listRes []itemvariantRes

	q := model.GetItemvariantQuery().Where().
		Int64("itemvariant.item_id", "=", req.ItemID).
		StringLike("itemvariant.itemvariant_name", req.ItemvariantName).
		StringLike("itemvariant.itemvariant_description", req.ItemvariantDescription).
		IsNull("itemvariant.delete_dt")

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	// get total page
	list, err = model.GetItemvariantWhere(ctx, conn, q)
	if err != nil {
		return err, cnt, listRes
	}
	cnt = len(list)

	// get data
	if req.GetPage() < 1 {
		req.SetPage(1)
	}

	list = make([]model.PublicItemvariant, 0)
	if req.GetPage() > 1 {
		q.OffsetLimit((req.GetPage()-1)*req.GetLimit(), req.GetLimit())
	} else {
		q.OffsetLimit(0, req.GetLimit())
	}

	list, err = model.GetItemvariantWhere(ctx, conn, q)
	if err != nil {
		return err, cnt, listRes
	}

	listRes = make([]itemvariantRes, 0)
	for _, data := range list {
		res := itemvariantRes{
			ItemvariantID:          data.ItemvariantID,
			ItemID:                 data.ItemID,
			ItemvariantName:        data.ItemvariantName,
			ItemvariantDescription: data.ItemvariantDescription,
			Price:                  data.Price,
			PhotoUrl:               getPhotoUrl(ctx, conn, data.PhotoID),
			IsActive:               data.IsActive,
			CreateBy:               data.CreateBy,
			CreateDt:               data.CreateDt,
			UpdateBy:               data.UpdateBy,
			UpdateDt:               data.UpdateDt,
		}
		listRes = append(listRes, res)
	}

	return err, cnt, listRes
}
