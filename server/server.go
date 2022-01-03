package main

import (
	"fmt"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()

	recipes := TestConnection()
	fmt.Println(recipes)
}
