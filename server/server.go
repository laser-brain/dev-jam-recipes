package main

import (
	"net/http"
	"strconv"

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

func main() {
	godotenv.Load()

	router := gin.Default()
	router.GET("/recipes", getRecipes)

	router.Run("localhost:8080")
}
