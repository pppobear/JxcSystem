package pagination

import (
	"fmt"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/module/handler"
)

func SetPageNav(resp *handler.ListResponse, c *gin.Context) {
	if resp.Page <= 1 {
		resp.PrePage = ""
	} else {
		resp.PrePage = fmt.Sprintf(
			"%s%s?page=%d&page_size=%d",
			c.Request.URL.Host,
			c.Request.URL.Path,
			resp.Page-1,
			resp.PageSize,
		)
	}
	if resp.Page >= resp.TotalPage {
		resp.NextPage = ""
	} else {
		resp.NextPage = fmt.Sprintf(
			"%s%s?page=%d&page_size=%d",
			c.Request.URL.Host,
			c.Request.URL.Path,
			resp.Page+1,
			resp.PageSize,
		)
	}
}
