package db

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Conn() *mongo.Client{
	_ = godotenv.Load("./.env")
	uri := os.Getenv("uri")
	// db := os.Getenv("dbname")
	// getcoll := os.Getenv("collname")


	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	
	if err != nil{
		fmt.Println(err)
	}
	
	// coll := client.Database(db).Collection(getcoll)
	return client
}
