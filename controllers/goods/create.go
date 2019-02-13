package goods

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type CreateRequest struct {
	Name           string `json:"name" binding:"required" validate:"required,min=1,max=20"`
	Model          string `json:"model" binding:"required" validate:"required,min=1,max=16"`
	Specifications string `json:"specifications" validate:"min=0,max=16"`
	UnitName       string `json:"unit_name" binding:"required" validate:"required,min=1,max=5"`
	MaxInventory   uint   `json:"max_inventory" validate:"min=0"`
	MinInventory   uint   `json:"min_inventory" validate:"min=0"`
}

func (r *CreateRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

func Create(c *gin.Context) {
	var r CreateRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	goods := models.GoodsModel{
		Id:             ksuid.New().String(),
		Name:           r.Name,
		Model:          r.Model,
		Specifications: r.Specifications,
		UnitName:       r.UnitName,
		MaxInventory:   r.MaxInventory,
		MinInventory:   r.MinInventory,
	}

	if err := goods.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	inv := models.InventoryModel{
		Id:     goods.Id,
		Number: 0,
	}

	if err := inv.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, goods)
}
