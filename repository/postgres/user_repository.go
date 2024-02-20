package postgres

import (
	"context"
	"database/sql"
	"fmt"
	"log"

	"github.com/jrione/gin-crud/domain"
)

type postgreUserRepository struct {
	DBClient *sql.DB
}

func NewUserRepository(conn *sql.DB) domain.UserRepository {
	return &postgreUserRepository{
		DBClient: conn,
	}
}

func (p *postgreUserRepository) GetAll(ctx context.Context) (res []domain.User, err error) {
	query := `SELECT username,full_name,email,password,created_at,updated_at FROM "User"`
	rows, err := p.DBClient.QueryContext(ctx, query)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Print(err)
		}
	}()

	res = make([]domain.User, 0)
	for rows.Next() {
		t := domain.User{}
		err := rows.Scan(
			&t.Username,
			&t.FullName,
			&t.Email,
			&t.Password,
			&t.Created_at,
			&t.Updated_at,
		)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		res = append(res, t)
	}
	return

}

func (p *postgreUserRepository) GetByUsername(ctx context.Context, username string) (*domain.User, error) {
	query := fmt.Sprintf(`SELECT username,full_name,email,password,created_at,updated_at FROM "User" WHERE username='%s'`, username)
	rows := p.DBClient.QueryRow(query)

	res := &domain.User{}
	err := rows.Scan(
		&res.Username,
		&res.FullName,
		&res.Email,
		&res.Password,
		&res.Created_at,
		&res.Updated_at,
	)

	if err != nil {
		log.Print(err)
		return nil, err
	}

	return res, nil

}
