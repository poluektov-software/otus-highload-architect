package main

import (
	"context"
	"log"
	"net/http"
	"otus-highload-architect/config"
	"otus-highload-architect/internal/model"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	_ "github.com/lib/pq"
	router "otus-highload-architect/internal/api"
	storageService "otus-highload-architect/internal/storage"
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
		httpServer    *http.Server
		storage       Storage
		routerService Router
		err           error
	)

	cfg, err := config.Get()
	if err != nil {
		log.Fatalf("Failed to get config: %s\n", err.Error())
	}

	conn, err = pgx.Connect(context.Background(), cfg.DatabaseURL)
	if err != nil {
		log.Fatalf("Unable to connect to database: %s\n", err.Error())
	}
	defer conn.Close(context.Background())

	storage = storageService.New(conn)
	routerService = router.New(storage)

	httpServer = &http.Server{
		Addr:           ":" + cfg.HostPort,
		Handler:        routerService.InitRoutes(),
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err = httpServer.ListenAndServe(); err != nil {
		log.Fatalf("Error occured while running http server: %s", err.Error())
	}
}
