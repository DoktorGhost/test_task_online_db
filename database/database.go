package database

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	_ "github.com/lib/pq"
)

var DB *sql.DB

/*
var Host = os.Getenv("DB_HOST")
var Port = os.Getenv("DB_PORT")
var User = os.Getenv("DB_USER")
var Password = os.Getenv("DB_PASSWORD")
var DBName = os.Getenv("DB_NAME")
*/

var Host = "localhost"
var Port = "5432"
var User = "admin"
var Password = "admin"
var DBName = "testdb"

// Инициализация базы данных
func InitDB() *sql.DB {
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s "+
		"password=%s dbname=%s sslmode=disable",
		Host, Port, User, Password, DBName)

	var DB *sql.DB
	var err error

	maxRetries := 5                  // Максимальное количество попыток соединения
	retryInterval := 5 * time.Second // Интервал между попытками

	for retries := 0; retries < maxRetries; retries++ {
		DB, err = sql.Open("postgres", psqlInfo)
		if err != nil {
			log.Printf("Failed to open database connection: %v", err)
			time.Sleep(retryInterval) // Подождите перед следующей попыткой
			continue
		}

		err = DB.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v", err)
			DB.Close() // Закройте соединение перед следующей попыткой
			time.Sleep(retryInterval)
			continue
		}

		return DB
	}

	log.Printf("Exhausted all connection retries, giving up.")
	return nil
}

func ExecuteSchemaSQL(db *sql.DB) {

	schemaSQL, err := ioutil.ReadFile("database/schema.sql")
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Exec(string(schemaSQL))
	if err != nil {
		log.Fatal(err)
	}
}
