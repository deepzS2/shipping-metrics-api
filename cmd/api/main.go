package main

import (
	"log"

	"github.com/deepzS2/shipping-metrics-api/cmd/docs"
	"github.com/deepzS2/shipping-metrics-api/internal/handler"
	"github.com/deepzS2/shipping-metrics-api/internal/repository"
	"github.com/deepzS2/shipping-metrics-api/internal/service"
	"github.com/deepzS2/shipping-metrics-api/pkg/config"
	"github.com/deepzS2/shipping-metrics-api/pkg/database"
	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

//	@title			Frete Rápido Backend Challenge API
//	@version		1.0
//	@description	An API for calculating shipping quotes and retrieving metrics.
//	@termsOfService	http://swagger.io/terms/

// @contact.name	deepzS2
// @contact.url	https://github.com/deepzS2
func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v", err)
	}

	log.Printf("%s", cfg.DSN())

	db, err := database.Connect(cfg.DSN(), "migrations")
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}

	defer db.Close()

	router := gin.Default()

	router.GET("/health", func(c *gin.Context) {
		c.JSON(200, gin.H{"status": "UP"})
	})

	freteRapidoService := service.NewFreteRapidoService(cfg.FreteRapidoAPIUrl)
	quoteRepository := repository.NewQuoteRepository(db)
	quoteService := service.NewQuoteService(freteRapidoService, quoteRepository)
	quoteHandler := handler.NewQuoteHandler(quoteService)

	apiGroup := router.Group("/api")
	{
		apiGroup.POST("/quote", quoteHandler.CreateQuote)
		apiGroup.GET("/metrics", quoteHandler.GetMetrics)
	}

	// Swagger
	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	log.Printf("Starting server on port %s", cfg.APIPort)

	if err := router.Run(":" + cfg.APIPort); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
