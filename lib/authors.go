package lib

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

type DBAuthor struct {
	Id         *int       `json:"id" db:"id"`
	FirstName  *string    `json:"firstname" db:"fname"`
	MiddleName *string    `json:"middlename" db:"mname"`
	LastName   *string    `json:"lastname" db:"lname"`
	Title      *string    `json:"title" db:"title"`
	IsActive   *bool      `json:"isactive" db:"isactive"`
	CreatedOn  *time.Time `json:"createdon" db:"createdOn"`
	UpdatedOn  *time.Time `json:"updatedon" db:"updatedOn"`
	DeletedOn  *time.Time `json:"deletedon" db:"deletedOn"`
}

func GetAuthors() ([]DBAuthor, error) {

	conn, err := GetPostgresConn("ie2")

	if err != nil {
		return nil, err
	}

	defer CloseConn(conn)

	qry := "SELECT * FROM authors"

	rows, err := conn.Query(context.Background(), qry)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBAuthor])
	if err != nil {
		return nil, err
	}

	return authors, nil
}

func GetAuthorsByName(name string) ([]DBAuthor, error) {

	conn, err := GetPostgresConn("ie2")

	if err != nil {
		return nil, err
	}

	defer CloseConn(conn)

	qry := `SELECT * 
			FROM authors 
			WHERE fname ILIKE '%' || $1 || '%'
           	   OR mname ILIKE '%' || $1 || '%'
               OR lname ILIKE '%' || $1 || '%'
	`

	rows, err := conn.Query(context.Background(), qry, name)

	if err != nil {
		return nil, err
	}

	defer rows.Close()

	authors, err := pgx.CollectRows(rows, pgx.RowToStructByName[DBAuthor])
	if err != nil {
		return nil, err
	}

	return authors, nil
}
