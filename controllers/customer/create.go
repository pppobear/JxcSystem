package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type CreateRequest struct {
	Name    string `json:"name" binding:"required" validate:"min=1,max=20"`
	Contact string `json:"contact" binding:"required" validate:"min=1,max=8"`
	Phone   string `json:"phone" binding:"required" validate:"min=1,max=20"`
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

	customer := models.CustomerModel{
		Id:      ksuid.New().String(),
		Name:    r.Name,
		Contact: r.Contact,
		Phone:   r.Phone,
	}

	if err := customer.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, customer)
}
