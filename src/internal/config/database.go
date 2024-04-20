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

	dbInfo := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s sslmode=disable port=%d",
		os.Getenv("DATABASE_HOST"),
		os.Getenv("DATABASE_USER"),
		os.Getenv("DATABASE_PASSWORD"),
		os.Getenv("DATABASE_NAME"),
		port,
	)
	database, err := sql.Open("postgres", dbInfo)

	if err != nil {
		panic(err)
	}

	err = database.Ping()
	if err != nil {
		panic(err)
	}

	return database, nil
}
