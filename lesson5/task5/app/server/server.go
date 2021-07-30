package server

import (
	"app/models"
	"app/repository/database"
	"app/routers"
	"app/usecase"
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct {
	httpServer *http.Server
	service    usecase.Usecase
}

func NewApp() *App {
	db := initDB()

	proxy := database.NewPostges(db)

	return &App{
		service: usecase.NewDBService(proxy),
	}
}

func (app *App) Run(port string) error {

	router := gin.Default()
	router.Use(
		gin.Recovery(),
		gin.Logger(),
	)

	api := router.Group("/api/v2")

	routers.RegisterNodeEndpoints(api, app.service)
	routers.RegisterDatabaseEndpoints(api, app.service)

	app.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	go func() {
		if err := app.httpServer.ListenAndServe(); err != nil {
			log.Fatalf("Failed to listen and serve: %+v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, os.Interrupt, os.Interrupt)

	<-quit

	ctx, shutdown := context.WithTimeout(context.Background(), 5*time.Second)
	defer shutdown()

	return app.httpServer.Shutdown(ctx)
}

func initDB() *gorm.DB {

	db, err := gorm.Open(postgres.Open(viper.GetString("db.dsn")), &gorm.Config{})
	if err != nil {
		log.Fatal("Error occured while establishing connection to Postgres")
	}

	db.AutoMigrate(&models.Nodes{})

	return db
}
