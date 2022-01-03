package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	page := 1
	var size int64 = 2
	var skip int64 = int64(page-1) * size
	recipesResponse := make(chan []Recipe)
	go GetAllRecipesPaged(&skip, &size, recipesResponse)
	recipes := <-recipesResponse
	fmt.Println(recipes)
}
