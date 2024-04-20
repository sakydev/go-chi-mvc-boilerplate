package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
)

func GetDatabase() (*sql.DB, error) {
	port, err := strconv.Atoi(os.Getenv("DATABASE_PORT"))
	if err != nil {
		panic("Error converting DATABASE_PORT to integer")
	}

	databaseOptions := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=%s port=%d",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		os.Getenv("DATABASE_SSL_MODE"),
		port,
	)
	database, err := sql.Open("postgres", databaseOptions)

	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	return database, nil
}
