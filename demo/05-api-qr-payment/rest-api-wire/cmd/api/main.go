package main

import (
	"context"
	"errors"
	"example/apiwire/cmd/api/di"
	"example/apiwire/cmd/api/router"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=example dbname=workshop port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("can not connect to database")
	}
	handlerWire := di.InitializeAPI(conn)
	ginSrv := router.GenerateRouter(handlerWire)

	srv := &http.Server{
		Addr:    ":8080",
		Handler: ginSrv,
	}

	go func() {
		logrus.Infof("[%s] http listen: %v", "qr-api", srv.Addr)

		if err := srv.ListenAndServe(); err != nil && errors.Is(err, http.ErrServerClosed) {
			logrus.Error("server listen err: ", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	logrus.Warn("shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		logrus.Fatal("server forced to shutdown: ", err)
	}

	logrus.Warn("server exited")
}
