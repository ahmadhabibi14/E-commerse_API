package conf

import (
	"database/sql"
	"fmt"
	"os"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

func ConnectDB() *sql.DB {
	zlog := InitLogger()
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
