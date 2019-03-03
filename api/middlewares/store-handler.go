package middlewares

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mateusmaaia/showcase-api/domains"
)

func ValidateStore() gin.HandlerFunc {
	return func(c *gin.Context) {
		store := c.Param("store")

		if len(store) == 0 {
			c.AbortWithStatusJSON(400, gin.H{"Error:": "Attribute store not set"})
			return
		}
		switch store {
		case domains.StoreName.Zap:
		case domains.StoreName.VivaReal:
			c.Set("store", store)
			c.Next()
			break
		default:
			c.AbortWithStatusJSON(400, gin.H{"Error:": fmt.Sprintf("Invalid value for attribute store [%s]", store)})
		}

	}
}
