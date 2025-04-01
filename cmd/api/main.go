package main

import (
	"context"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.uber.org/fx"

	"events-api/internal/config"
	"events-api/internal/handlers"
	"events-api/internal/middleware"
	"events-api/internal/repositories"
	"events-api/internal/services"
	"events-api/pkg/database"

	_ "events-api/docs" // Importa la documentación generada

	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title			Events API
// @version		1.0
// @description	API para gestión de eventos
// @host			localhost:8080
// @BasePath		/api/v1
func main() {
	app := fx.New(
		// Proporciona todas las dependencias
		fx.Provide(
			config.NewConfig,
			database.NewMongoClient,
			repositories.NewEventRepository,
			services.NewEventService,
			handlers.NewEventHandler,
			newGinRouter,
		),
		// Registra los hooks del ciclo de vida
		fx.Invoke(registerHooks),
	)

	// Inicia la aplicación
	startCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Start(startCtx); err != nil {
		log.Fatal(err)
	}

	// Espera señales para un apagado elegante
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	<-c

	// Detiene la aplicación
	stopCtx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	if err := app.Stop(stopCtx); err != nil {
		log.Fatal(err)
	}
}

// Crea una nueva instancia del router Gin
func newGinRouter() *gin.Engine {
	r := gin.Default()
	r.Use(middleware.Logger())

	// Rutas Swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return r
}

// Registra las rutas HTTP y otras configuraciones
func registerHooks(
	lc fx.Lifecycle,
	router *gin.Engine,
	eventHandler *handlers.EventHandler,
	mongoClient *mongo.Client,
	cfg *config.Config,
) {
	lc.Append(fx.Hook{
		OnStart: func(context.Context) error {
			// Configuración de rutas
			v1 := router.Group("/api/v1")
			{
				events := v1.Group("/events")
				{
					events.POST("", eventHandler.CreateEvent)
					events.GET("", eventHandler.GetAllEvents)
					events.GET("/:id", eventHandler.GetEventByID)
					events.PUT("/:id", eventHandler.UpdateEvent)
					events.DELETE("/:id", eventHandler.DeleteEvent)
					events.PUT("/:id/review", eventHandler.ReviewEvent)
					events.PUT("/:id/unreview", eventHandler.UnreviewEvent)
					events.GET("/types", eventHandler.GetEventTypes)
					events.POST("/seed", eventHandler.SeedEvents)
					events.GET("/status", eventHandler.GetEventStatus)
					events.GET("/management-status", eventHandler.GetEventManagementStatus)
					events.GET("/management-required", eventHandler.GetEventsRequiringManagement)
					events.GET("/no-management-required", eventHandler.GetEventsNotRequiringManagement)
				}
			}

			// Inicia el servidor HTTP
			go func() {
				if err := router.Run(":" + cfg.Port); err != nil {
					log.Printf("Error al iniciar el servidor: %v\n", err)
				}
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {
			log.Println("Desconectando de MongoDB...")
			return mongoClient.Disconnect(ctx)
		},
	})
}
