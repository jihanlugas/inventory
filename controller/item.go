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

type Item struct{}

func ItemComposer() Item {
	return Item{}
}

type pageItemReq struct {
	request.Paging
	ItemName        string `json:"itemName" form:"itemName" query:"itemName"`
	ItemDescription string `json:"itemDescription" form:"itemDescription" query:"itemDescription"`
}

type createItemReq struct {
	ItemName        string                `json:"itemName" form:"itemName" validate:"required,lte=200"`
	ItemDescription string                `json:"itemDescription" form:"itemDescription" validate:"lte=200"`
	IsActive        bool                  `json:"isActive" form:"isActive" validate:""`
	Photo           *multipart.FileHeader `json:"-"`
	PhotoChk        bool                  `json:"photo" validate:"required,photo=Photo"`
}

type itemRes struct {
	ItemID          int64      `json:"itemId"`
	PropertyID      int64      `json:"propertyId"`
	ItemName        string     `json:"itemName"`
	ItemDescription string     `json:"itemDescription"`
	PhotoUrl        string     `json:"photoUrl"`
	IsActive        bool       `json:"isActive"`
	CreateBy        int64      `json:"createBy"`
	CreateDt        *time.Time `json:"createDt"`
	UpdateBy        int64      `json:"updateBy"`
	UpdateDt        *time.Time `json:"updateDt"`
}

// Page Page Item
// @Summary Page Item
// @Tags Item
// @Accept json
// @Produce json
// @Param req body pageItemReq true "payload"
// @Success      200  {object}	response.Response{payload=response.Pagination}
// @Failure      500  {object}  response.Response
// @Router /item [post]
func (h Item) Page(c echo.Context) error {
	var err error

	req := new(pageItemReq)
	if err = c.Bind(req); err != nil {
		errorInternal(c, err)
	}

	if err = c.Validate(req); err != nil {
		return response.StatusBadRequest("validation error", response.ValidationError(err)).SendJSON(c)
	}

	err, cnt, list := getPageItem(req)
	if err != nil {
		errorInternal(c, err)
	}

	return response.StatusOk("success", response.PayloadPagination(req, list, cnt)).SendJSON(c)
}

// Create Create Item
// @Summary Create Item
// @Tags Item
// @Accept json
// @Produce json
// @Param itemName formData string true "Item Name"
// @Param itemDescription formData string true "Item Description"
// @Param isActive formData boolean true "Active"
// @Param photo formData file true "Photo"
// @Success      200  {object}	response.Response{payload=itemRes}
// @Failure      500  {object}  response.Response
// @Router /item/create [post]
func (h Item) Create(c echo.Context) error {
	var err error
	var data model.PublicItem
	var publicphoto model.PublicPhoto

	loginUser, err := getUserLoginInfo(c)
	if err != nil {
		errorInternal(c, err)
	}

	req := new(createItemReq)
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

	data.PropertyID = loginUser.PropertyID
	data.ItemName = req.ItemName
	data.ItemDescription = req.ItemDescription
	data.IsActive = req.IsActive
	data.CreateBy = loginUser.UserID
	data.UpdateBy = loginUser.UserID
	err = data.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	publicphoto.CreateBy = loginUser.UserID
	err = publicphoto.Upload(conn, ctx, tx, req.Photo, constant.PhotoRefTableItem)
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

	res := itemRes{
		ItemID:          data.ItemID,
		PropertyID:      data.PropertyID,
		ItemName:        data.ItemName,
		ItemDescription: data.ItemDescription,
		PhotoUrl:        getPhotoUrl(ctx, conn, data.PhotoID),
		IsActive:        data.IsActive,
		CreateBy:        data.CreateBy,
		CreateDt:        data.CreateDt,
		UpdateBy:        data.UpdateBy,
		UpdateDt:        data.UpdateDt,
	}

	return response.StatusCreated("success", res).SendJSON(c)
}

// GetById Get Item
// @Summary Get Item
// @Tags Item
// @Accept json
// @Produce json
// @Param kanji path number true "item_id" default(0)
// @Success      200  {object}	response.Response{payload=itemRes}
// @Failure      500  {object}  response.Response
// @Router /item/{item_id} [get]
func (h Item) GetById(c echo.Context) error {
	var err error
	var data model.PublicItem

	ID, err := strconv.ParseInt(c.Param("item_id"), 10, 64)
	if err != nil {
		return response.StatusNotFound("data not found", response.Payload{}).SendJSON(c)
	}

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	data.ItemID = ID
	err = data.GetById(ctx, conn)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			return response.StatusNotFound("data not found", response.Payload{}).SendJSON(c)
		}
		errorInternal(c, err)
	}

	var res itemRes
	res.ItemID = data.ItemID
	res.PropertyID = data.PropertyID
	res.ItemName = data.ItemName
	res.ItemDescription = data.ItemDescription
	res.PhotoUrl = getPhotoUrl(ctx, conn, data.PhotoID)
	res.IsActive = data.IsActive
	res.CreateBy = data.CreateBy
	res.CreateDt = data.CreateDt
	res.UpdateBy = data.UpdateBy
	res.UpdateDt = data.UpdateDt

	return response.StatusOk("success", res).SendJSON(c)
}

func getPageItem(req *pageItemReq) (error, int, []itemRes) {
	var err error
	var cnt int
	var list []model.PublicItem
	var listRes []itemRes

	q := model.GetItemQuery().Where().
		StringLike("item.item_name", req.ItemName).
		StringLike("item.item_description", req.ItemDescription)

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	// get total page
	list, err = model.GetItemWhere(ctx, conn, q)
	if err != nil {
		return err, cnt, listRes
	}
	cnt = len(list)

	// get data
	if req.GetPage() < 1 {
		req.SetPage(1)
	}

	list = make([]model.PublicItem, 0)
	if req.GetPage() > 1 {
		q.OffsetLimit((req.GetPage()-1)*req.GetLimit(), req.GetLimit())
	} else {
		q.OffsetLimit(0, req.GetLimit())
	}

	list, err = model.GetItemWhere(ctx, conn, q)
	if err != nil {
		return err, cnt, listRes
	}

	listRes = make([]itemRes, 0)
	for _, data := range list {
		res := itemRes{
			ItemID:          data.ItemID,
			PropertyID:      data.PropertyID,
			ItemName:        data.ItemName,
			ItemDescription: data.ItemDescription,
			PhotoUrl:        getPhotoUrl(ctx, conn, data.PhotoID),
			IsActive:        data.IsActive,
			CreateBy:        data.CreateBy,
			CreateDt:        data.CreateDt,
			UpdateBy:        data.UpdateBy,
			UpdateDt:        data.UpdateDt,
		}
		listRes = append(listRes, res)
	}

	return err, cnt, listRes
}
