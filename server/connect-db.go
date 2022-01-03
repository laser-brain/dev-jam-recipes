package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Recipe struct {
	Id          int    `json:"id" bson:"_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Thumbnail   string `json:"thumbnail"`
}

func TestInsert() {
	data := TestConnection()
	var id int
	if data == nil {
		id = 1
	} else {
		id = len(data) + 1
	}

	recipe := &Recipe{
		Id:          id,
		Name:        "test-db-recipe",
		Description: "This is data from mongo database",
		Thumbnail:   "https://www.themealdb.com/images/logo-small.png",
	}
	InsertRecipe(recipe)
}

func InsertRecipe(recipe *Recipe) {
	ctx, client, cancel := loadCollection()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("recipes").Collection("recipes")

	res, insertErr := collection.InsertOne(ctx, recipe)
	if insertErr != nil {
		log.Fatal(insertErr)
	}
	fmt.Printf("Insert response: %q\n", res)
}

func GetAllRecipesPaged(page int, size int) {

}

func TestConnection() []Recipe {
	ctx, client, cancel := loadCollection()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("recipes").Collection("recipes")

	cur, currErr := collection.Find(ctx, bson.D{})
	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var recipes []Recipe
	err := cur.All(ctx, &recipes)
	if err != nil {
		panic(err)
	}

	return recipes
}

func loadCollection() (ctx context.Context, client *mongo.Client, Cancel context.CancelFunc) {
	dbUser := os.Getenv("MONGO_DB_USER")
	dbPassword := os.Getenv("MONGO_DB_PASSWORD")
	dbUrl := os.Getenv("MONGO_DB_URL")
	connectionString := "mongodb+srv://" + dbUser + ":" + dbPassword + "@" + dbUrl + "?retryWrites=true&w=majority"
	client, err := mongo.NewClient(
		options.Client().ApplyURI(connectionString),
	)
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)

	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	return ctx, client, cancel
}
