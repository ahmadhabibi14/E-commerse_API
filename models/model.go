package models

import (
	"e-commerse_api/conf"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

var (
	MySQLTables = []string{ // Define SQL tables here
		`Product`,
	}
	migrationsDir = `file://db/migrations`
)

func RunMigration() {
	conf.LoadEnv()
	zlog := conf.InitLogger()
	db := conf.ConnectDB()
	defer db.Close()

	driver, _ := mysql.WithInstance(db, &mysql.Config{})
	m, err := migrate.NewWithDatabaseInstance(migrationsDir, os.Getenv(`MYSQL_NAME`), driver)
	if err != nil {
		zlog.Fatal().Msg(`Error: ` + err.Error())
	}
	// Run the migration
	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		zlog.Fatal().Msg(`Error: ` + err.Error())
	}

	for _, table := range MySQLTables {
		query := fmt.Sprintf("SHOW CREATE TABLE %s", table)
		rows, err := db.Query(query)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
		defer rows.Close()
		var tableNameResult, createTableStatement string
		for rows.Next() {
			err := rows.Scan(&tableNameResult, &createTableStatement)
			if err != nil {
				zlog.Fatal().Msg(`Error: ` + err.Error())
			}
		}
		sqlSchemaFile := fmt.Sprintf("db/schema/%s.sql", table)
		file, err := os.Create(sqlSchemaFile)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
		defer file.Close()
		_, err = file.WriteString(createTableStatement)
		if err != nil {
			zlog.Fatal().Msg(`Error: ` + err.Error())
		}
	}

	fmt.Println("Migration successful !!")
}
