package drivers

import (
	"fmt"
	"net/http"
	"strings"
	"time"

	jwtlib "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	"github.com/gin-gonic/gin"
	"github.com/json-iterator/go"

	"pppobear.cn/jxc-backend/config"
	"pppobear.cn/jxc-backend/connections/redis"
)

var json = jsoniter.ConfigCompatibleWithStandardLibrary

type jwtAuthManager struct {
	secret string
	exp    time.Duration
	alg    string
}

func NewJwtAuthDriver() *jwtAuthManager {
	return &jwtAuthManager{
		secret: config.GetJwtConfig().Secret,
		exp:    config.GetJwtConfig().Exp,
		alg:    config.GetJwtConfig().Alg,
	}
}

// Check the token of request header is valid or not.
func (jwtAuth *jwtAuthManager) Check(c *gin.Context) bool {
	token := c.Request.Header.Get("Authorization")
	token = strings.Replace(token, "Bearer ", "", -1)
	if token == "" {
		return false
	}
	if result, _ := redis.Client.Get(token); result != "1" {
		return false
	}
	var keyFun jwtlib.Keyfunc
	keyFun = func(token *jwtlib.Token) (interface{}, error) {
		b := []byte(jwtAuth.secret)
		return b, nil
	}
	authJwtToken, err := request.ParseFromRequest(c.Request, request.OAuth2Extractor, keyFun)

	if err != nil {
		fmt.Println(err)
		return false
	}

	c.Set("User", map[string]interface{}{
		"token": authJwtToken,
	})

	return authJwtToken.Valid
}

// User is get the auth user from token string of the request header which
// contains the user ID. The token string must start with "Bearer "
func (jwtAuth *jwtAuthManager) User(c *gin.Context) interface{} {

	var jwtToken *jwtlib.Token
	if jwtUser, exist := c.Get("User"); !exist {
		tokenStr := c.Request.Header.Get("Authorization")
		tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
		if tokenStr == "" {
			return map[interface{}]interface{}{}
		}
		var err error
		jwtToken, err = jwtlib.Parse(tokenStr, func(token *jwtlib.Token) (interface{}, error) {
			b := []byte(jwtAuth.secret)
			return b, nil
		})
		if err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
	} else {
		jwtToken = jwtUser.(map[string]interface{})["token"].(*jwtlib.Token)
	}

	if claims, ok := jwtToken.Claims.(jwtlib.MapClaims); ok && jwtToken.Valid {
		var user map[string]interface{}
		if err := json.Unmarshal([]byte(claims["user"].(string)), &user); err != nil {
			fmt.Println(err)
			return map[interface{}]interface{}{}
		}
		c.Set("User", map[string]interface{}{
			"token": jwtToken,
			"user":  user,
		})
		return user
	} else {
		fmt.Println(ok)
		return map[interface{}]interface{}{}
	}
}

func (jwtAuth *jwtAuthManager) Login(http *http.Request, w http.ResponseWriter, user map[string]interface{}) interface{} {
	token := jwtlib.New(jwtlib.GetSigningMethod(jwtAuth.alg))
	// Set some claims
	userStr, err := json.Marshal(user)
	token.Claims = jwtlib.MapClaims{
		"user": string(userStr),
		"exp":  time.Now().Add(jwtAuth.exp).Unix(),
	}
	// Sign and get the complete encoded token as a string
	tokenString, err := token.SignedString([]byte(jwtAuth.secret))
	if err != nil {
		return nil
	}
	redis.Client.Set(tokenString, 1, time.Hour*24*3)
	return tokenString
}

func (jwtAuth *jwtAuthManager) Logout(c *gin.Context) bool {
	if !jwtAuth.Check(c) {
		return false
	}
	tokenStr := c.Request.Header.Get("Authorization")
	tokenStr = strings.Replace(tokenStr, "Bearer ", "", -1)
	redis.Client.Del(tokenStr)
	return true
}
