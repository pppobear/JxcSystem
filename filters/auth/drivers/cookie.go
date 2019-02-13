package drivers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"

	"pppobear.cn/jxc-backend/config"
)

var store = sessions.NewCookieStore([]byte(config.GetEnv().AppSecret))

type cookieAuthManager struct {
	name string
}

func NewCookieAuthDriver() *cookieAuthManager {
	return &cookieAuthManager{
		name: config.GetCookieConfig().Name,
	}
}

func (cookie *cookieAuthManager) Check(c *gin.Context) bool {
	// read cookie
	session, err := store.Get(c.Request, cookie.name)
	return err == nil &&
		session != nil &&
		session.Values != nil &&
		session.Values["id"] != nil
}

func (cookie *cookieAuthManager) User(c *gin.Context) interface{} {
	// get model user
	session, _ := store.Get(c.Request, cookie.name)
	return session.Values
}

func (cookie *cookieAuthManager) Login(
	http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	// write cookie
	session, err := store.Get(http, cookie.name)
	if err != nil {
		return false
	}
	session.Values["id"] = user["id"]
	_ = session.Save(http, w)
	return true
}

func (cookie *cookieAuthManager) Logout(c *gin.Context) bool {
	// del cookie
	session, err := store.Get(c.Request, cookie.name)
	if err != nil {
		return false
	}
	session.Values["id"] = nil
	_ = session.Save(c.Request, c.Writer)
	return true
}
