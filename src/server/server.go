package server

import (
	"time"

	loginRouters "LogiredAPIWeb/src/core/services/auth/infrastructure"
	userRouters "LogiredAPIWeb/src/internal/users/infrastructure"

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	"LogiredAPIWeb/src/server/middleware"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Run(
	authRoutes *loginRouters.AuthRoutes,
	userRoutes *userRouters.UserRoutes,
) {

	r := gin.Default()
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	globalLimiter := middleware.NewRateLimiter(20, 40)
	r.Use(globalLimiter.Middleware())

	r.Use(middleware.MaxBodySize(20 * 1024 * 1024))

	r.Use(middleware.Timeout(30 * time.Second))

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Content-Type", "Authorization"},
		AllowCredentials: false,
	}))

	authGroup := r.Group("/")
	authGroup.Use(middleware.StrictRateLimit(5.0/60.0, 5))
	authRoutes.AttachRoutes(r)

	userRoutes.AttachRoutes(r)

	r.Run(":8081")
}
