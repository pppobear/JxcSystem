package staff

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/segmentio/ksuid"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

type CreateStaffRequest struct {
	Name         *string           `json:"name" binding:"required" validate:"min=1,max=8"`
	Gender       *string           `json:"gender" validate:"max=2"`
	Birthday     *handler.JsonDate `json:"birthday"`
	SpecialtyId  *uint64           `json:"specialty_id"`
	DepartmentId *uint64           `json:"department_id"`
	Married      *bool             `json:"married"`
}

func (r *CreateStaffRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

type CreateSpeRequest struct {
	Name string `json:"name" binding:"required" validate:"min=1,max=50"`
}

func (r *CreateSpeRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

type CreateDepRequest struct {
	Name     string `json:"name" binding:"required" validate:"min=1,max=50"`
	HeadName string `json:"head_name" binding:"required" validate:"min=1,max=50"`
}

func (r *CreateDepRequest) Validate() error {
	return handler.ValidateAdapt(r)
}

func CreateStaff(c *gin.Context) {
	var r CreateStaffRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	staff := models.StaffModel{
		Id:           ksuid.New().String(),
		Name:         *r.Name,
		Gender:       *r.Gender,
		Birthday:     *r.Birthday,
		SpecialtyId:  *r.SpecialtyId,
		DepartmentId: *r.DepartmentId,
	}
	if r.Married != nil {
		staff.Married = sql.NullBool{Bool: *r.Married, Valid: true}
	}

	if err := staff.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, staff)
}

func CreateSpe(c *gin.Context) {
	var r CreateSpeRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	spe := models.SpecialtyModel{
		Name: r.Name,
	}

	if err := spe.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, spe)
}

func CreateDep(c *gin.Context) {
	var r CreateDepRequest
	if err := c.Bind(&r); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	// Validate the data.
	if err := r.Validate(); err != nil {
		handler.SendResponse(c, errno.ErrValidation, nil)
		return
	}

	dep := models.DepartmentModel{
		Name:     r.Name,
		HeadName: r.HeadName,
	}

	if err := dep.Create(); err != nil {
		handler.SendResponse(c, errno.ErrDatabase, nil)
		return
	}

	handler.SendResponse(c, nil, dep)
}
