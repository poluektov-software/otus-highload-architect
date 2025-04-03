package storage

import (
	"cmd/internal/model"
	"context"
	"fmt"
	"github.com/jackc/pgx/v5"
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
		fmt.Printf("SQL insert failed: %v\n", err)
	}
	return
}

func (s *service) GetUser(ctx context.Context, id int) (user model.Account, err error) {
	err = s.db.QueryRow(ctx, "select id, first_name, last_name from public.accounts where id=$1", id).Scan(&user.Id, &user.FirstName, &user.LastName)
	if err != nil {
		fmt.Printf("SQL select failed: %v\n", err)
	}
	return
}

func New(db *pgx.Conn) Service {
	return &service{
		db: db,
	}
}
