package database

import (
	"database/sql"
	"fmt"

	"github.com/gobuffalo/packr/v2"
	migrate "github.com/rubenv/sql-migrate"
)

var (
	DBConnection *sql.DB
)

func DbMigrate(dbParam *sql.DB) {
	migrations := &migrate.PackrMigrationSource{
		Box: packr.New("migrations", "./sql_migrations"),
	}

	n, errs := migrate.Exec(dbParam, "postgres", migrations, migrate.Up)
	if errs != nil {
		panic(errs)
	}

	DBConnection = dbParam

	fmt.Println("Applied", n, " migrations!")
}
