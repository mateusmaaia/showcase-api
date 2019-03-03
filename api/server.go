package api

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/mateusmaaia/showcase-api/api/middlewares"
	"github.com/mateusmaaia/showcase-api/controllers"
)

type Server struct {
	RealEstateController controllers.RealEstateController
	HealthController     controllers.HealthController
}

func (s *Server) Run() {
	router := gin.Default()
	router.GET("real-estates/:store", middlewares.ValidateStore(), middlewares.ValidateParams(), s.RealEstateController.Get)
	router.GET("health", s.HealthController.Check)
	router.Run(os.Getenv("PORT"))
}
