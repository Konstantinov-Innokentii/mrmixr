package postgres

import (
	"fmt"
	"github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

type Postgres struct {
	DB      *sqlx.DB
	Builder squirrel.StatementBuilderType
}

func New(dbname, user, password string) (*Postgres, error) {
	pg := &Postgres{}
	pg.Builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Dollar)
	pgConnectionString := fmt.Sprintf("user=%s dbname=%s password=%s sslmode=disable", user, dbname, password)
	db, err := sqlx.Connect("postgres", pgConnectionString)
	if err != nil {
		return nil, err
	}
	pg.DB = db
	return pg, nil
}
