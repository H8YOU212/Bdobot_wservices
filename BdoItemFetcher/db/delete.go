package db

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
)

func Del(client mongo.Client, filter interface{}) {
	getdb, getcoll := GetENV()

	cl := client.Database(getdb).Collection(getcoll)

	result, err := cl.DeleteOne(context.TODO(), filter)

	if err != nil{
		fmt.Println(err)
	} else {
		fmt.Println(result)
	}

}