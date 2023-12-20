package conf

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var zlog = InitLogger()

func init() {
	// Used for testing
	if os.Getenv("WEB_ENV") == "dev" {
		err := godotenv.Load("../.env")
		if err != nil {
			zlog.Error().
				Str("ERROR", err.Error()).
				Msg("cannot load .env files")
		}
	}
}

func ConnectDB() *sql.DB {
	DbDriver := "mysql"
	DbHost := os.Getenv("MYSQL_HOST")
	DbPort := os.Getenv("MYSQL_PORT")
	DbName := os.Getenv("MYSQL_NAME")
	DbUser := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open(DbDriver, dsn)
	if err != nil {
		zlog.Panic().
			Str("ERROR", err.Error()).
			Msg("cannot connect to " + DbDriver)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)
	return db
}
