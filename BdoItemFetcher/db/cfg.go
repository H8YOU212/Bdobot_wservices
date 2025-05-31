package db

import (
	"os"

	"github.com/joho/godotenv"
)

func GetENV() (string, string) {
	godotenv.Load("./.env")
	

	db := os.Getenv("dbname")
	coll := os.Getenv("collname")

	return db, coll
}