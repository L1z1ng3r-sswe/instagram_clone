package app

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/db"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/handler"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/repository"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/internal/service"
	"github.com/L1z1ng3r-sswe/instagram_clone/app/pkg/logging"
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
)

type App struct {
	Handler *handler.Handler
}

func NewApp() *App {
	a := &App{}

	logger := logging.GetLogger()

	if err := initConfigs(); err != nil {
		errMsg := "Failed to confige init: " + err.Error()
		logger.Ftl(errMsg)
	}

	db, err := getDB()
	if err != nil {
		errMsg := "Failed to init db: " + err.Error()
		logger.Ftl(errMsg)
	}

	handler := initDep(db, logger)
	a.Handler = handler

	return a
}

func initConfigs() error {
	if err := godotenv.Load(".env"); err != nil {
		return err
	}

	return nil
}

func getDB() (*sqlx.DB, error) {
	dbConfig := db.DBConfig{
		Host:     os.Getenv("DATABASE_HOST"),
		Port:     os.Getenv("DATABASE_PORT"),
		Password: os.Getenv("DATABASE_DB_PASSWORD"),
		User:     os.Getenv("DATABASE_USER_NAME"),
		DBName:   os.Getenv("DATABASE_DB_NAME"),
	}

	return db.InitDb(dbConfig)
}

func initDep(db *sqlx.DB, log *logging.Logger) *handler.Handler {
	r := repository.NewRepository(db)
	s := service.NewService(r)
	h := handler.NewHandler(s, log)

	return h
}

func (a *App) Run(port string, handler *gin.Engine) error {
	server := &http.Server{
		Addr:              ":" + port,
		Handler:           handler,
		MaxHeaderBytes:    20 << 20,
		WriteTimeout:      10 * time.Second,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()

		if err := server.Shutdown(ctx); err != nil {
			fmt.Printf("Server shutdown error: %s\n", err.Error())
		}
	}()

	if err := server.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
