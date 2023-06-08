package main

import (
	"context"
	"errors"
	userRepo "example/restapi/internal/repository/user"
	"example/restapi/internal/services/user"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Handler struct {
	UserService user.UserService
}

func (h Handler) RegisterUserHandler(c *gin.Context) {
	if err := h.UserService.CreateUser(c.Request.Context(), "superAdmin"); err != nil {
		logrus.WithError(err).Error("cannot register user")
		c.JSON(http.StatusInternalServerError, gin.H{"msg": "internal error"})
	}
	c.JSON(http.StatusCreated, gin.H{"Msg": "hi"})
}

func generateRouter(h Handler) *gin.Engine {
	router := gin.New()
	router.Use(cors.Default())
	router.Use(gin.Recovery())

	router.GET("/user", h.RegisterUserHandler)

	return router
}

func main() {
	conn, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  "host=localhost user=postgres password=example dbname=workshop port=5432 sslmode=disable TimeZone=Asia/Bangkok",
		PreferSimpleProtocol: true, // disables implicit prepared statement usage
	}), &gorm.Config{})
	if err != nil {
		panic("can not connect to database")
	}

	userRepository := userRepo.ProvideUserRepo(conn)
	userService := user.ProvideUserService(userRepository)
	handler := Handler{
		UserService: userService,
	}
	ginSrv := generateRouter(handler)

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
