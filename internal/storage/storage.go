package storage

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"

	"otus-highload-architect/internal/model"
)

type Service interface {
	AddUser(ctx context.Context, user model.Account) (err error)
	GetUser(ctx context.Context, id int) (user model.Account, err error)
}

type service struct {
	db *pgx.Conn
}

func (s *service) AddUser(ctx context.Context, user model.Account) (err error) {
	_, err = s.db.Exec(ctx, "insert into public.accounts (first_name, last_name) values ($1, $2)", user.FirstName, user.LastName)
	if err != nil {
		log.Printf("SQL insert failed: %s\n", err.Error())
	}
	return
}

func (s *service) GetUser(ctx context.Context, id int) (user model.Account, err error) {
	err = s.db.QueryRow(ctx, "select id, first_name, last_name from public.accounts where id=$1", id).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		log.Printf("SQL select failed: %s\n", err.Error())
	}
	return
}

func New(db *pgx.Conn) Service {
	return &service{
		db: db,
	}
}
