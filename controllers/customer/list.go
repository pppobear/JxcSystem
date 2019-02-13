package customer

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
	"pppobear.cn/jxc-backend/module/pagination"
)

func List(c *gin.Context) {
	var req handler.ListRequest
	if err := c.MustBindWith(&req, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}

	resp, err := models.ListCustomer(&req)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}
