package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picbed/controllers"
	"picbed/utils"
)

func AuthMiddleware(c *gin.Context) {
	authHeader := c.Request.Header.Get("Authorization")
	if authHeader == "" {
		c.JSON(http.StatusUnauthorized, controllers.ResponseData{
			Code: controllers.CodeUnauthorized,
			Msg:  controllers.CodeUnauthorized.Msg(),
			Data: nil,
		})
		c.Abort()
		return
	}
	token := authHeader[len("Bearer "):]
	claims, err := utils.ParseJWTToken(token)
	if err != nil {
		c.JSON(http.StatusUnauthorized, controllers.ResponseData{
			Code: controllers.CodeUnauthorized,
			Msg:  controllers.CodeUnauthorized.Msg(),
			Data: nil,
		})
		c.Abort()
		return
	}
	c.Set("userInfo", claims)
	c.Next()
}
