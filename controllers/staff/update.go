package staff

import (
	"strconv"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type UpdateStaffRequest struct {
	Name     *string           `json:"name"`
	Gender   *string           `json:"gender" validate:"max=2"`
	Birthday *handler.JsonDate `json:"birthday"`
	Married  *bool             `json:"married"  validate:"max=1"`
}

func (r *UpdateStaffRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

type UpdateSpeRequest struct {
	Name string `json:"name" validate:"min=1"`
}

func (r *UpdateSpeRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

type UpdateDepRequest struct {
	Name     string `json:"name"`
	HeadName string `json:"head_name"`
}

func (r *UpdateDepRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

func UpdateStaff(c *gin.Context) {
	var r UpdateStaffRequest
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

	staff := models.StaffModel{
		Id: id,
	}

	if err := staff.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, staff)
}

func UpdateSpe(c *gin.Context) {
	var r UpdateSpeRequest
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
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

	spe := models.SpecialtyModel{
		Id:   id,
		Name: r.Name,
	}
	if err := spe.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, spe)
}

func UpdateDep(c *gin.Context) {
	var r UpdateDepRequest
	id, err := strconv.ParseUint(c.Param("id"), 10, 64)
	if err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	dep := models.DepartmentModel{
		Id:       id,
		Name:     r.Name,
		HeadName: r.HeadName,
	}
	if err := dep.Update(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}
	handler.SendResponse(c, nil, dep)
}
