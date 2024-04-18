package main

import (
	"fmt"
	"github.com/ClickHouse/clickhouse-go/v2"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog/log"
	"github.com/swaggo/files"
	"github.com/swaggo/gin-swagger"
	"stat4market_task/config"
	"stat4market_task/docs"
	"stat4market_task/handler"
	"stat4market_task/repository"
	"stat4market_task/service"
	"time"
)

func main() {
	cfg := config.GetConfig()

	conn, err := clickhouse.Open(&clickhouse.Options{
		Addr: []string{cfg.Clickhouse.Url},
		Auth: clickhouse.Auth{
			Database: "test",
			Username: "user1",
			Password: "123456",
		},
		//Debug:           true,
		DialTimeout:     time.Second,
		MaxOpenConns:    10,
		MaxIdleConns:    5,
		ConnMaxLifetime: time.Hour,
	})
	if err != nil {
		log.Error().Msgf("error connecting to clickhouse : %v", err)
	}

	eventRepo := repository.NewEventRepo(conn)
	eventService := service.EventService(eventRepo)
	eventHandler := handler.EventHandler(eventService)

	router := gin.Default()
	docs.SwaggerInfo.BasePath = "/"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	rEvent := router.Group("/api")
	{
		rEvent.POST("/event", eventHandler.Post)
	}

	err = router.Run(fmt.Sprintf("%s:%s", cfg.Router.ServerAddr, cfg.Router.ServerPort))
	if err != nil {
		log.Error().Msgf("error running gin-router: %v", err)
	}
}
