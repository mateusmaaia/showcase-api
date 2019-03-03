package middlewares

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ValidateParams() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var pageSize int
		var pageNumber int
		pageSize, _ = strconv.Atoi(ctx.Query("pageSize"))
		pageNumber, _ = strconv.Atoi(ctx.Query("pageNumber"))

		if pageSize == 0 {
			pageSize = 50
		}

		if pageNumber == 0 {
			pageNumber = 1
		}

		ctx.Set("pageSize", pageSize)
		ctx.Set("pageNumber", pageNumber)
		ctx.Next()
	}
}
