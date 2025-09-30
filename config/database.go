package config

import (
	"context"
	"database/sql"
	"fmt"

	otelsql "github.com/XSAM/otelsql"
	_ "github.com/lib/pq"
	semconv "go.opentelemetry.io/otel/semconv/v1.4.0"
)

func OpenDB(env *Config) *sql.DB {

	conn := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v",
		env.Database.Username,
		env.Database.Password,
		env.Database.Hostname,
		env.Database.Port,
		env.Database.Name,
		env.Database.SSLMode,
	)

	client, err := otelsql.Open(
		"postgres",
		conn,
		otelsql.WithAttributes(semconv.DBSystemPostgreSQL),
		otelsql.WithSpanNameFormatter(func(ctx context.Context, method otelsql.Method, query string) string {
			return env.OTLP.SqlPrefix + string(method) + ": " + query
		}),

		otelsql.WithSQLCommenter(true),
	)

	if err != nil {
		return nil
	}
	if err = client.Ping(); err != nil {
		panic(err)
	}

	err = otelsql.RegisterDBStatsMetrics(client, otelsql.WithAttributes(
		semconv.DBSystemPostgreSQL,
	))
	if err != nil {
		panic(err)
	}

	return client
}
