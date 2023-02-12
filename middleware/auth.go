package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/net-http/database"
	"github.com/net-http/models"
	"github.com/net-http/utils"
)

func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokens := strings.Split(c.GetHeader("Authorization"), " ")
		if len(tokens) != 2 {
			utils.CustomRepsonseWriter(c, http.StatusUnauthorized, nil, "please provide bearer token")
			c.Abort()
			return
		}
		tokenStr := tokens[1]

		//decoding mail from token
		email := utils.DecodeBasicAuthToken(tokenStr)

		//get the user from the database
		var user models.User
		if err := database.Db.Where("email =?", email).First(&user).Error; err != nil {
			utils.CustomRepsonseWriter(c, http.StatusNotFound, nil, "invalid token")
			c.Abort()
			return
		}
		c.Next()
	}
}
