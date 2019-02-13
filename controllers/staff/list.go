package staff

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
	"pppobear.cn/jxc-backend/module/pagination"
)

type LSRequest struct {
	handler.ListRequest
	models.StaffFilterRequest
}

func ListStaff(c *gin.Context) {
	var (
		lsr LSRequest
		req handler.ListRequest
		ft  models.StaffFilterRequest
	)
	if err := c.MustBindWith(&lsr, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	req = lsr.ListRequest
	ft = lsr.StaffFilterRequest

	resp, err := models.ListStaff(&req, &ft)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}

func ListSpe(c *gin.Context) {
	var req handler.ListRequest
	if err := c.MustBindWith(&req, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := models.ListSpecialty(&req)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}

func ListDep(c *gin.Context) {
	var req handler.ListRequest
	if err := c.MustBindWith(&req, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := models.ListDepartment(&req)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}
