package controller

import (
	"github.com/jihanlugas/inventory/db"
	"github.com/jihanlugas/inventory/model"
	"github.com/jihanlugas/inventory/request"
	"github.com/jihanlugas/inventory/response"
	"github.com/labstack/echo/v4"
)

type Item struct{}

func ItemComposer() Item {
	return Item{}
}

type pageItemReq struct {
	request.Paging
}

type createItemReq struct {
	PropertyID      int64  `db:"property_id,use_zero" json:"propertyId" form:"propertyId" validate:"required"`
	ItemName        string `db:"item_name,use_zero" json:"itemName" form:"itemName" validate:"required,lte=200"`
	ItemDescription string `db:"item_description,use_zero" json:"itemDescription" form:"itemDescription" validate:"required,lte=200"`
	Price           int64  `db:"price,use_zero" json:"price" form:"price" validate:"required"`
	IsActive        bool   `db:"is_active,use_zero" json:"isActive" form:"isActive" validate:"required"`
}

// Page Page Item
// @Summary Page Item
// @Tags Item
// @Accept json
// @Produce json
// @Param req query pageItemReq true "json req body"
// @Success      200  {object}	response.Response{payload=response.Pagination}
// @Failure      500  {object}  response.Response
// @Router /item [get]
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
// @Param req body createItemReq true "json req body"
// @Success      200  {object}	response.Response{payload=model.ItemRes}
// @Failure      500  {object}  response.Response
// @Router /item/create [post]
func (h Item) Create(c echo.Context) error {
	var err error
	var data model.PublicItem

	loginUser, err := getUserLoginInfo(c)
	if err != nil {
		errorInternal(c, err)
	}

	req := new(createItemReq)
	if err = c.Bind(req); err != nil {
		return err
	}

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
	data.Price = req.Price
	data.IsActive = req.IsActive
	data.CreateBy = loginUser.UserID
	data.UpdateBy = loginUser.UserID
	err = data.Insert(ctx, tx)
	if err != nil {
		errorInternal(c, err)
	}

	if err = tx.Commit(ctx); err != nil {
		_ = tx.Rollback(ctx)
		errorInternal(c, err)
	}

	return response.StatusCreated("success", data.Res()).SendJSON(c)
}

func getPageItem(req *pageItemReq) (error, int, []model.PublicItem) {
	var err error
	var cnt int
	var list []model.PublicItem

	q := model.GetItemQuery().Where()

	conn, ctx, closeConn := db.GetConnection()
	defer closeConn()

	// get total page
	list, err = model.GetItemWhere(ctx, conn, q)
	if err != nil {
		return err, cnt, list
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
		return err, cnt, list
	}

	return err, cnt, list
}
