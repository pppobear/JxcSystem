package purchases

import (
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"

	"pppobear.cn/jxc-backend/filters/auth"
	"pppobear.cn/jxc-backend/models"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
	"pppobear.cn/jxc-backend/module/pagination"
)

type LPRequest struct {
	handler.ListRequest
	models.PurchasesFilterRequest
}

// @Summary 获取 销售记录
// @Description 列出 所有的销售记录。若用户不是管理员，则仅显示自己所产生的记录。
// @Tags 销售
// @Security JWT
// @Accept  json
// @Produce  json
// @Param page path string false "当前的页数" default(1)
// @Param page_size path string false "每页显示记录的数目" default(id)
// @Param order_by path string false "排序的根据, 在开头添加-表示为" default(id) Enums(id, datetime, C)
// @Param start_date path string false "要查询的起始时间，默认为所有"
// @Param end_date path string false "要查询的截止时间，默认为所有"
// @Param search path string false "模糊搜索条件，支持：单号、供应商id、供应商名称、销售人id、销售人名称 字段的搜索"
// @Success 200 {object} handler.Response
// @Failure 200 {object} handler.Response
// @Router /purchases [get]
func List(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	user := (*authDr).User(c).(map[string]interface{})
	var (
		lpr LPRequest
		req handler.ListRequest
		ft  models.PurchasesFilterRequest
	)
	if err := c.MustBindWith(&lpr, binding.Form); err != nil {
		handler.SendResponse(c, errno.ErrBind, nil)
		return
	}
	req = lpr.ListRequest
	ft = lpr.PurchasesFilterRequest

	resp, err := models.ListPurchases(&req, &ft, &user)
	if err == nil {
		pagination.SetPageNav(&resp, c)
	}

	handler.SendResponse(c, err, resp)
}
