package utils

import (
	"blog-backend/global"
	"blog-backend/model/system"
	systemReq "blog-backend/model/system/request"
	"net"
	"time"

	"github.com/gin-gonic/gin"
)

func ClearToken(c *gin.Context) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", "", -1, "/", "", false, false)
	} else {
		c.SetCookie("x-token", "", -1, "/", host, false, false)
	}
}

func SetToken(c *gin.Context, token string, maxAge int) {
	host, _, err := net.SplitHostPort(c.Request.Host)
	if err != nil {
		host = c.Request.Host
	}

	if net.ParseIP(host) != nil {
		c.SetCookie("x-token", token, maxAge, "/", "", false, false)
	} else {
		c.SetCookie("x-token", token, maxAge, "/", host, false, false)
	}
}

func GetToken(c *gin.Context) string {
	token, _ := c.Cookie("x-token")
	if token == "" {
		j := NewJWT()
		token = c.Request.Header.Get("x-token")
		claims, err := j.ParseToken(token)
		if err != nil {
			global.YAGAMI_LOGGER.Error("重新写入cookie token失败,未能成功解析token,请检查请求头是否存在x-token且claims是否为规定结构")
			return token
		}
		SetToken(c, token, int((claims.ExpiresAt.Unix()-time.Now().Unix())/60))
	}
	return token
}

func GetClaims(c *gin.Context) (*systemReq.CustomClaims, error) {
	token := GetToken(c)
	j := NewJWT()
	claims, err := j.ParseToken(token)
	if err != nil {
		global.YAGAMI_LOGGER.Error("从Gin的Context中获取从jwt解析信息失败, 请检查请求头是否存在x-token且claims是否为规定结构")
	}
	return claims, err
}

func GetUserId(c *gin.Context) uint {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return 0
		} else {
			return cl.BaseClaims.Id
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.Id
	}
}

func GetUserName(c *gin.Context) string {
	if claims, exists := c.Get("claims"); !exists {
		if cl, err := GetClaims(c); err != nil {
			return ""
		} else {
			return cl.Name
		}
	} else {
		waitUse := claims.(*systemReq.CustomClaims)
		return waitUse.BaseClaims.Name
	}
}

func LoginToken(user system.Login) (token string, claims systemReq.CustomClaims, err error) {
	j := &JWT{SigningKey: []byte(global.YAGAMI_CONFIG.JWT.SigningKey)}
	claims = j.CreateClaims(systemReq.BaseClaims{
		Id:   user.GetUserId(),
		Name: user.GetUsername(),
	})
	token, err = j.CreateToken(claims)
	if err != nil {
		return
	}
	return
}
