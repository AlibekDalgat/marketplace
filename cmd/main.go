package main

import (
	"context"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
	"marketplace/internal/app"
	"marketplace/internal/config"
	"marketplace/internal/delivery"
	"marketplace/internal/repository"
	"marketplace/internal/service"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	logrus.SetFormatter(new(logrus.JSONFormatter))

	dbCfg := config.GetDBConfig()
	db, err := repository.OpenDB(dbCfg)
	if err != nil {
		logrus.Fatalf("Ошибка при подклюении к базе данных: %s", err.Error())
	}

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := delivery.NewHandler(services)

	srv := new(app.Server)
	go func() {
		if err := srv.Run(os.Getenv("HTTP_PORT"), handlers.InitRoutes()); err != nil && err != http.ErrServerClosed {
			logrus.Fatalf("Ошибка при работе http-сервера: %s", err.Error())
		}
	}()

	logrus.Println("Маркетплейс начал работы")
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGTERM, syscall.SIGINT)
	<-quit

	logrus.Println("Маркетплейс завершил работу")
	if err := srv.Shutdown(context.Background()); err != nil {
		logrus.Errorf("Произошла ошибка при завершении работы сервера: %s", err.Error())
	}
	if err := db.Close(); err != nil {
		logrus.Errorf("Ошибка при отсоединении от базы данных: %s", err.Error())
	}
}
