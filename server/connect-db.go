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
	Id          int          `json:"id" bson:"_id"`
	Name        string       `json:"name"`
	Type        string       `json:"type"`
	Difficulty  string       `json:"difficulty"`
	Description string       `json:"description"`
	Thumbnail   string       `json:"thumbnail"`
	Servings    int          `json:"servings"`
	Ingredients []Ingredient `json:"ingredients"`
}

type Ingredient struct {
	Name   string `json:"name"`
	Amount string `json:"amount"`
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

func GetAllRecipesPaged(skip *int64, size *int64, responseChannel chan []Recipe) {

	ctx, client, cancel := loadCollection()
	defer client.Disconnect(ctx)
	defer cancel()

	collection := client.Database("recipes").Collection("recipes")

	opts := &options.FindOptions{
		Skip:  skip,
		Limit: size,
	}
	cur, currErr := collection.Find(ctx, bson.D{}, opts)

	if currErr != nil {
		panic(currErr)
	}
	defer cur.Close(ctx)

	var recipes []Recipe
	err := cur.All(ctx, &recipes)
	if err != nil {
		panic(err)
	}

	responseChannel <- recipes
}

func loadCollection() (context.Context, *mongo.Client, context.CancelFunc) {
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
