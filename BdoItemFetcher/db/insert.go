package db

import (
	"context"
	"fmt"
	"time"

	// "go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Insert(client *mongo.Client, doc []interface{}) {
	getdb, getcoll := GetENV() 
	
	cl := client.Database(getdb).Collection(getcoll)

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()


	result, err := cl.InsertMany(ctx, doc, options.InsertMany())
	if err != nil{
		fmt.Println("")
	} else {
		fmt.Println(result)
	}

}
