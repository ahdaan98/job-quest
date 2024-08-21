package middleware

import (
	"github.com/ahdaan67/JobQuest/pkg/helper"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func AdminAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		tokenHeader := c.GetHeader("Authorization") // Capitalize header key

		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenpart := splitted[1]
		tokenClaims, err := helper.ValidateToken(tokenpart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Check if role is admin (this assumes ValidateToken returns a role field)
		if tokenClaims.Role != "admin" {
			response := response.ClientResponse(http.StatusForbidden, "Insufficient permissions", nil, nil)
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		c.Set("tokenClaims", tokenClaims)

		c.Next()
	}
}
