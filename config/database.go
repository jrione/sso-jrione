package config

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func OpenDB(env *Config) *sql.DB {

	conn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?%v", env.Database.Username, env.Database.Password, env.Database.Hostname, env.Database.Port, env.Database.Name, "sslmode=disable")
	client, err := sql.Open("postgres", conn)
	if err != nil {
		return nil
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}

	return client
}
