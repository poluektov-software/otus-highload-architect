package api

import (
	"cmd/internal/model"
	"context"
	"github.com/gin-gonic/gin"
)

type Storage interface {
	AddUser(ctx context.Context, user model.Account) (err error)
	GetUser(ctx context.Context, id int) (user model.Account, err error)
}

type Router interface {
	InitRoutes() *gin.Engine
}

type service struct {
	storage Storage
}

func (s *service) InitRoutes() *gin.Engine {
	router := gin.New()
	//router.POST("/login", h.login)
	router.POST("/user/register", s.create)
	router.GET("/user/get/:id", s.get)
	return router
}

func New(storage Storage) Router {
	return &service{
		storage: storage,
	}
}
