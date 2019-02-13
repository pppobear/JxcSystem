package user

import (
	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

func Logout(c *gin.Context) {
	authDr, _ := c.MustGet("jwt_auth").(*auth.Auth)
	if (*authDr).Logout(c) {
		handler.SendResponse(c, nil, map[string]interface{}{
			"message": "Logout successfully. Token would be abundant."})
		return
	}
	handler.SendResponse(c, errno.ErrLogoutFailed, nil)
}
