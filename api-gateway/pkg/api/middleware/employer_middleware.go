package middleware

import (
	"github.com/ahdaan67/JobQuest/pkg/helper"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func EmployerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization") // Ensure the correct capitalization

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

		tokenPart := splitted[1]
		tokenClaims, err := helper.ValidateTokenEmployer(tokenPart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Optional: Check if the role is correct
		if tokenClaims.Role != "employer" {
			response := response.ClientResponse(http.StatusUnauthorized, "Unauthorized role", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		// Store the claims in the context if needed later
		c.Set("id", tokenClaims.Id)
		c.Set("role", tokenClaims.Role)

		c.Next()
	}
}