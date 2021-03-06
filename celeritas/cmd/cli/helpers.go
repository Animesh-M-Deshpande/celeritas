package main

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/joho/godotenv"
)

func setup() {

	err := godotenv.Load()

	if err != nil {
		exitGracefully(err)
	}

	path, err := os.Getwd()

	if err != nil {
		exitGracefully(err)
	}

	cel.RootPath = path
	cel.DB.DataType = os.Getenv("DATABASE_TYPE")
}
func getDSN() string {
	dbType := cel.DB.DataType

	if dbType == "pgx" {
		dbType = "postgres"
	}

	if dbType == "postgres" {
		var dsn string

		if os.Getenv("DATABASE_PASS") != "" {
			dsn = fmt.Sprintf("postgres://%s:%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_PASS"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		} else {
			dsn = fmt.Sprintf("postgres://%s@%s:%s/%s?sslmode=%s",
				os.Getenv("DATABASE_USER"),
				os.Getenv("DATABASE_HOST"),
				os.Getenv("DATABASE_PORT"),
				os.Getenv("DATABASE_NAME"),
				os.Getenv("DATABASE_SSL_MODE"))
		}

		return dsn
	}
	return "mysql://" + cel.BuildDSN()

}

func showHelp() {
	color.Yellow(`Available commands:

	help           			- show the help commands
	version        			- print application version
	migrate		   			- runs all up migrations that have not been previously
	migrate down   			- recereses the most recent migrations
	migrate reset  			- runs all down migrations in rever order, and then all up migrations
	make migration <name>	 - creates to new up and down migrations in the migrations folder
	make auth				- Create and run migrations for authentication table and create models and middlewares
	make handler <name>		- Creates a stub handler in the handler directory
	make models <name>		- creates a new model in the data directory
	make session			- creates a table in the database as a session store
	`)
}
