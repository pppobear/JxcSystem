package goods

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type UpdateRequest struct {
	Name           string `json:"name" validate:"min=0,max=20"`
	Model          string `json:"model" validate:"min=0,max=16"`
	Specifications string `json:"specifications" validate:"min=0,max=16"`
	UnitName       string `json:"unit_name" validate:"min=0,max=5"`
	MaxInventory   uint   `json:"max_inventory" validate:"min=0"`
	MinInventory   uint   `json:"min_inventory" validate:"min=0"`
}

func (r *UpdateRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

type UpdateInvRequest struct {
	Number    uint    `json:"number" validate:"min=0"`
	UnitPrice float64 `json:"unit_price" validate:"min=0"`
}

func (r *UpdateInvRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

func UpdateGood(c *gin.Context) {
	var r UpdateRequest
	id := c.Param("id")
	if len(id) == 0 {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	goods := models.GoodsModel{
		Id:             id,
		Name:           r.Name,
		Model:          r.Model,
		Specifications: r.Specifications,
		UnitName:       r.UnitName,
		MaxInventory:   r.MaxInventory,
		MinInventory:   r.MinInventory,
	}

	if err := goods.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, goods)
}

func UpdateInventory(c *gin.Context) {
	var r UpdateInvRequest
	id := c.Param("id")
	if len(id) == 0 {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	inv := models.InventoryModel{
		Id:        id,
		Number:    r.Number,
		UnitPrice: r.UnitPrice,
	}
	if err := inv.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, inv)
}
