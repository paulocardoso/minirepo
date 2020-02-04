package v1

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
)


type User struct {
	ID string `json:"id"`
}

const (
	URL = "mongodb://localhost:27017"
)

func GetUser(c *gin.Context) {


	GetCollection("users")
	collection.InsertOne(c , User{"1"})

	var results []*User

	cur, err := collection.Find(context.TODO(), bson.D{})

	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem User
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		results = append(results, &elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())


	c.JSON(http.StatusOK, results)
}


func GetCollection(name string) * mongo.Collection{
	//Base Context
	ctx := context.Background()

	// Options to the database.
	clientOpts := options.Client().ApplyURI(URL)
	client, err := mongo.Connect(ctx, clientOpts)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	log.Println("Connected to MongoDB!")

	return client.Database("test").Collection(name)

}