package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mateusmaaia/showcase-api/api/middlewares"
	"github.com/mateusmaaia/showcase-api/controllers"
)

type Server struct {
	realEstateController controllers.RealEstateController
	healthController     controllers.HealthController
}

func (s *Server) Run() {
	router := gin.Default()
	router.GET("real-estates/:store", middlewares.ValidateStore(), middlewares.ValidateParams(), s.realEstateController.Get)
	router.GET("health", s.healthController.Check)
	router.Run(":8090")
}
