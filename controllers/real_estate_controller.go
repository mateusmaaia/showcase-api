package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mateusmaaia/showcase-api/api/responses"
	"github.com/mateusmaaia/showcase-api/services"
)

type RealEstateController struct {
	RealEstateService services.RealEstateService
}

func (r *RealEstateController) Get(ctx *gin.Context) {
	store := ctx.Param("store")
	pageSize := ctx.GetInt("pageSize")
	pageNumber := ctx.GetInt("pageNumber")
	realEstates, end := r.RealEstateService.FindByStore(store, pageSize, pageNumber)
	total := r.RealEstateService.CountByStore(store)

	realEstatesResponse := &responses.RealEstateResponse{PageSize: end,
		PageNumber: pageNumber,
		TotalCount: total,
		Listings:   realEstates,
	}
	ctx.JSON(http.StatusOK, realEstatesResponse)
}
