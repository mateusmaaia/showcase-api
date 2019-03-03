package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mateusmaaia/showcase-api/api/responses"
	"github.com/mateusmaaia/showcase-api/services"
)

type RealEstateController struct {
	realEstateService services.RealEstateService
}

func (r *RealEstateController) Get(ctx *gin.Context) {
	store := ctx.Param("store")
	pageSize, _ = strconv.Atoi(ctx.Param("pageSize"))
	pageNumber, _ = strconv.Atoi(ctx.Param("pageNumber"))
	realEstates := r.realEstateService.FindByStore(store, pageSize, pageSize*pageNumber)
	total := r.realEstateService.CountByStore(store)

	realEstatesResponse := &responses.RealEstateResponse{PageSize: pageSize,
		PageNumber: pageNumber,
		TotalCount: total,
		Listings:   realEstates,
	}
	ctx.JSON(http.StatusOK, realEstatesResponse)
}
