package conf

import (
	"time"

	"../jwt"

	"../models"
	"github.com/gin-gonic/gin"
)

var JWTMiddlewareConf *jwt.GinJWTMiddleware

func InitJWTMiddlewareConf() *jwt.GinJWTMiddleware {

	JWTMiddlewareConf = &jwt.GinJWTMiddleware{
		Realm:         "test zone",
		Key:           []byte("secret key"),
		Timeout:       time.Hour,
		MaxRefresh:    time.Hour,
		Authenticator: models.AuthenticateUser,
		Authorizator: func(email string, c *gin.Context) bool {
			// if email == "rotblauer@gmail.com" {
			// 	return true
			// }
			// return false
			return true
		},
		Unauthorized: func(c *gin.Context, code int, message string) {
			c.JSON(code, gin.H{
				"code":    code,
				"message": message,
			})
		},
		PayloadFunc: func(userId string) map[string]interface{} {
			m := make(map[string]interface{})
			m["email"] = userId
			return m
		},
	}

	return JWTMiddlewareConf
}
