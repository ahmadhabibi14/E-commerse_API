package test

import (
	"database/sql"
	"fmt"
	"os"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func TestConnectDB(t *testing.T) {
	err := godotenv.Load("../.env")
	if err != nil {
		t.Log(`Error: `, err)
	}
	DbDriver := "mysql"
	DbHost := os.Getenv("MYSQL_HOST")
	DbPort := os.Getenv("MYSQL_PORT")
	DbName := os.Getenv("MYSQL_NAME")
	DbUser := os.Getenv("MYSQL_USER")
	DbPassword := os.Getenv("MYSQL_PASSWORD")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true", DbUser, DbPassword, DbHost, DbPort, DbName)
	db, err := sql.Open(DbDriver, dsn)
	if err != nil {
		t.Log(`Error connect DB: `, err)
	}

	db.SetMaxIdleConns(10)
	db.SetMaxOpenConns(100)
	db.SetConnMaxIdleTime(5 * time.Minute)
	db.SetConnMaxLifetime(60 * time.Minute)

	t.Log(`Success connect to database`)
}
