package sales

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"pppobear.cn/jxc-backend/filters/auth"
	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
	"pppobear.cn/jxc-backend/module/pagination"
)

type LSRequest struct {
	handler.ListRequest
	models.SalesFilterRequest
}

func List(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	user := (*authDr).User(c).(map[string]interface{})
	var (
		lsr LSRequest
		req handler.ListRequest
		ft  models.SalesFilterRequest
	)
	if err := c.MustBindWith(&lsr, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	req = lsr.ListRequest
	ft = lsr.SalesFilterRequest

	resp, err := models.ListSales(&req, &ft, &user)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}
