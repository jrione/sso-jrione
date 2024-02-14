package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/jrione/gin-crud/domain"
)

type postgreStudentRepository struct {
	DBClient *sql.DB
}

func NewStudentRepository(conn *sql.DB) domain.StudentRepository { // define connection
	return &postgreStudentRepository{conn}
}

func (p *postgreStudentRepository) fetch(ctx context.Context, query string) (result []domain.Student, err error) {
	rows, err := p.DBClient.QueryContext(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	defer func() {
		errRow := rows.Close()
		if errRow != nil {
			log.Fatal(err)
		}
	}()

	res := make([]domain.Student, 0)
	for rows.Next() {
		t := domain.Student{}
		err := rows.Scan(
			&t.ID,
			&t.Name,
			&t.Age,
			&t.Grade,
		)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		res = append(res, t)
	}
	return res, nil
}

func (p *postgreStudentRepository) Fetch(ctx context.Context) (res []domain.Student, err error) {
	query := `SELECT*FROM tb_student`

	res, err = p.fetch(ctx, query)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	return

}
