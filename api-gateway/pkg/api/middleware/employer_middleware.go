package middleware

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/ahdaan67/JobQuest/pkg/helper"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"

	"github.com/gin-gonic/gin"
)

func EmployerAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenHeader := c.GetHeader("Authorization")

		if tokenHeader == "" {
			response := response.ClientResponse(http.StatusUnauthorized, "No auth header provided", nil, nil)
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		splitted := strings.Split(tokenHeader, " ")
		if len(splitted) != 2 || splitted[0] != "Bearer" {
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

		fmt.Printf("Token Claims: %+v\n", tokenClaims)

		if tokenClaims.Role != "employer" {
			response := response.ClientResponse(http.StatusForbidden, "Forbidden: Insufficient Role", nil, nil)
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		c.Set("id", tokenClaims.Id)
        c.Set("role", tokenClaims.Role)

		c.Next()
	}
}
