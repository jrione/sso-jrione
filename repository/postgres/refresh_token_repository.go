package postgres

import (
	"context"
	"database/sql"
	"log"

	"github.com/jrione/gin-crud/domain"
)

type postgreRefreshTokenRepository struct {
	DBClient *sql.DB
}

func NewRefreshTokenRepository(conn *sql.DB) domain.RefreshTokenRepository {
	return &postgreRefreshTokenRepository{
		DBClient: conn,
	}
}

func (p *postgreRefreshTokenRepository) StoreRefreshToken(ctx context.Context, data domain.RefreshTokenData) (err error) {
	query := `INSERT INTO refresh_token(username,refresh_token) VALUES( $1, $2 )`
	state, err := p.DBClient.PrepareContext(ctx, query)
	if err != nil {
		log.Printf("Error PrepareContext: %s", err)
		return err
	}
	defer state.Close()
	_, err = state.ExecContext(ctx, data.Username, data.RefreshToken)
	if err != nil {
		log.Printf("Error Exec: %s", err)
		return err
	}
	return nil
}

func (p *postgreRefreshTokenRepository) GetRefreshToken(ctx context.Context, us string) (res *domain.RefreshTokenData, err error) {
	query := `SELECT username,refresh_token FROM refresh_token WHERE username=$1`
	state, err := p.DBClient.PrepareContext(ctx, query)
	log.Print(query)
	if err != nil {
		log.Printf("Error PrepareContext: %s", err)
		return nil, err
	}
	defer state.Close()
	row := state.QueryRow(us)

	res = &domain.RefreshTokenData{}
	err = row.Scan(
		&res.Username,
		&res.RefreshToken,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		log.Print(err)
		return nil, err
	}
	return res, nil

}

// func (p *postgreRefreshTokenRepository) UpdateRefreshToken(ctx context.Context, refreshToken string) (string, error) {
// 	return "", nil
// }

// func (p *postgreRefreshTokenRepository) DeleteRefreshToken(ctx context.Context, refreshToken string) (ok bool, err error) {
// 	return false, nil
// }
