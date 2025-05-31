package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func UpdateColl(client *mongo.Client, doc []interface{}, filter interface{}) {
	getdb, getcoll := GetENV()
	cl := client.Database(getdb).Collection(getcoll)

	result, err := cl.UpdateMany(context.TODO(), filter, doc, options.Update())

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}
}
