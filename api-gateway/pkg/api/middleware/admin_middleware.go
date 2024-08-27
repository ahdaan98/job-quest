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
		// Retrieve the Authorization header
		tokenHeader := c.GetHeader("Authorization")

		// Check if the Authorization header is provided
		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Split the Authorization header to get the token
		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token Format", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		tokenPart := splitted[1]

		// Validate the token and extract claims
		tokenClaims, err := helper.ValidateToken(tokenPart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Check if the role is admin
		if tokenClaims.Role != "admin" {
			response := response.ClientResponse(http.StatusForbidden, "Insufficient permissions", nil, nil)
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		// Store token claims in the context
		c.Set("tokenClaims", tokenClaims)

		// Proceed to the next handler
		c.Next()
	}
}
