package middleware

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/ahdaan67/JobQuest/pkg/helper"
	"github.com/ahdaan67/JobQuest/pkg/utils/response"
)

func JobSeekerAuthMiddleware() gin.HandlerFunc {
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
		tokenClaims, err := helper.ValidateTokenJobSeeker(tokenPart)
		if err != nil {
			response := response.ClientResponse(http.StatusUnauthorized, "Invalid Token", nil, err.Error())
			c.JSON(http.StatusUnauthorized, response)
			c.Abort()
			return
		}

		if tokenClaims.Role != "jobseeker" {
			response := response.ClientResponse(http.StatusForbidden, "Forbidden: Insufficient Role", nil, nil)
			c.JSON(http.StatusForbidden, response)
			c.Abort()
			return
		}

		jobseekerID := int32(tokenClaims.Id)
		c.Set("id", jobseekerID)

		c.Next()
	}
}
