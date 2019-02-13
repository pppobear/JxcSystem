package customer

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type UpdateRequest struct {
	Name    string `json:"name" binding:"required" validate:"min=1,max=20"`
	Contact string `json:"contact" binding:"required" validate:"min=1,max=8"`
	Phone   string `json:"phone" binding:"required" validate:"min=1,max=20"`
}

func (r *UpdateRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

func Update(c *gin.Context) {
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

	customer := models.CustomerModel{
		Id:      id,
		Name:    r.Name,
		Contact: r.Contact,
		Phone:   r.Phone,
	}

	if err := customer.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, customer)
}
