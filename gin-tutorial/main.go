package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type manga struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"artist"`
	Price  float64 `json:"price"`
}

var mangas = []manga{
	{ID: 1, Title: "Berserk", Author: "Kentaro Miura", Price: 25.50},
	{ID: 2, Title: "BLue Lock", Author: "some chick", Price: 32.30},
	{ID: 3, Title: "Hunter x Hunter", Author: "forgot his name", Price: 44.40},
	{ID: 4, Title: "JJBA", Author: "Araki", Price: 64.54},
	{ID: 5, Title: "Parasyte", Author: "don't know", Price: 54.25},
}

func getMangas(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mangas)
}

func main() {
	router := gin.Default()
	router.GET("/mangas", getMangas)
	router.Run("localhost:8080")
}
