package config

import "database/sql"

type App struct {
	Env      *Config
	DBClient *sql.DB
}

func Bootstrap() App {
	app := &App{}
	app.Env = NewEnv()
	app.DBClient = OpenDB(app.Env)

	return *app
}
