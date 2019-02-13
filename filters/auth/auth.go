package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"pppobear.cn/jxc-backend/filters/auth/drivers"
	"pppobear.cn/jxc-backend/module/errno"
	"pppobear.cn/jxc-backend/module/handler"
)

var driverList = map[string]func() Auth{
	"cookie": func() Auth {
		return drivers.NewCookieAuthDriver()
	},
	"jwt": func() Auth {
		return drivers.NewJwtAuthDriver()
	},
}

type Auth interface {
	Check(c *gin.Context) bool
	User(c *gin.Context) interface{}
	Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{}
	Logout(c *gin.Context) bool
}

func RegisterGlobalAuthDriver(authKey string, key string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		c.Set(key, driver)
		c.Next()
	}
}

func Middleware(authKey string) gin.HandlerFunc {
	return func(c *gin.Context) {
		driver := GenerateAuthDriver(authKey)
		if !(*driver).Check(c) {
			handler.SendResponse(c, errno.ErrTokenAuthFailed, nil)
			c.Abort()
		}
		c.Next()
	}
}

func GenerateAuthDriver(string string) *Auth {
	var authDriver Auth
	authDriver = driverList[string]()
	return &authDriver
}

func GetCurUser(c *gin.Context, key string) map[string]interface{} {
	authDriver, _ := c.MustGet(key).(*Auth)
	return (*authDriver).User(c).(map[string]interface{})
}

func User(c *gin.Context) map[string]interface{} {
	return GetCurUser(c, "jwt_auth")
}
