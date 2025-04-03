package main

import (
	"cmd/http_server"
	"cmd/internal/model"
	"context"
	"fmt"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/sirupsen/logrus"

	router "cmd/internal/api"
	storageService "cmd/internal/storage"
	_ "github.com/lib/pq"
)

type Storage interface {
	AddUser(ctx context.Context, user model.Account) (err error)
	GetUser(ctx context.Context, id int) (user model.Account, err error)
}

type Router interface {
	InitRoutes() *gin.Engine
}

func main() {
	var (
		conn          *pgx.Conn
		storage       Storage
		routerService Router
		err           error
	)

	conn, err = pgx.Connect(context.Background(), "postgres://otus:otus@localhost:5434/otus")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		os.Exit(1)
	}
	defer conn.Close(context.Background())

	storage = storageService.New(conn)
	routerService = router.New(storage)

	srv := new(http_server.Server)
	if err = srv.Run("8098", routerService.InitRoutes()); err != nil {
		logrus.Fatalf("error occured while running http server: %s", err.Error())
	}
}
