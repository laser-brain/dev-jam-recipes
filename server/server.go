package main

import (
	"flag"
	"net/http"
	"os"
	"strconv"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func getRecipes(c *gin.Context) {
	pageParam := c.Query("page")
	page, err := strconv.ParseInt(pageParam, 10, 32)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}
	sizeParam := c.Query("size")
	size, err := strconv.ParseInt(sizeParam, 10, 64)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	skip := (page - 1) * size
	recipesResponse := make(chan []Recipe)
	go GetAllRecipesPaged(&skip, &size, recipesResponse)
	recipes := <-recipesResponse
	c.JSON(http.StatusOK, gin.H{"recipes": recipes})
}

func addRecipe(c *gin.Context) {
	var recipe Recipe
	if err := c.BindJSON(&recipe); err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
	}
	InsertRecipe(&recipe)
}

func seedDatabase() {
	for i := 1; i < 21; i++ {
		stringIndex := strconv.Itoa(i)
		recipe := &Recipe{
			Id:          i,
			Name:        "Test recipe " + stringIndex,
			Type:        "Supper",
			Difficulty:  "Beginner",
			Description: "This is the description of the recipe #" + stringIndex + " from the cloud-database.",
			Servings:    5,
			Thumbnail:   "https://www.themealdb.com/images/logo-small.png",
			Ingredients: []Ingredient{{Name: "Meat", Amount: "much"}},
		}

		InsertRecipe(recipe)
	}
}

func main() {
	godotenv.Load()
	cmd := flag.String("cmd", "", "")
	flag.Parse()

	cmdValue := string(*cmd)
	if cmdValue == "seed" {
		seedDatabase()
	}

	router := gin.Default()
	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "PUT", "POST", "DELETE", "PATCH"},
		AllowHeaders:     []string{"*"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))
	router.GET("/recipes", getRecipes)
	router.POST("/recipes", addRecipe)

	port := os.Getenv("PORT")
	router.Run(":" + port)
}
