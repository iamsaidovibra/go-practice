package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type manga struct {
	ID     uint    `json:"id"`
	Title  string  `json:"title"`
	Author string  `json:"author"`
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

func postMangas(c *gin.Context) {
	var newManga manga
	fmt.Println()
	fmt.Println("newManga before:", newManga)
	fmt.Println()

	if err := c.BindJSON(&newManga); err != nil {
		return
	}

	fmt.Println()
	fmt.Println("newManga after:", newManga)
	fmt.Println()

	mangas = append(mangas, newManga)
	c.IndentedJSON(http.StatusAccepted, newManga)

	fmt.Println()
	fmt.Println("mangas:", mangas)
	fmt.Println()
}

func getMangaByID(c *gin.Context) {
	id := c.Param("id")
	for _, a := range mangas {
		convertedCounter := strconv.FormatUint(uint64(a.ID), 10)
		if convertedCounter == id {
			c.IndentedJSON(http.StatusOK, a)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "manga not found"})
}

func deleteManga(c *gin.Context) {
	id := c.Param("id")
	for _, a := range mangas {
		convertedCounter := strconv.FormatUint(uint64(a.ID), 10)
		if convertedCounter == id {
			sqlStatement := `DELETE FROM manga WHERE id = %d;`
			_, err := db.Exec(sqlStatement, 1)
			if err != nil {
				panic(err)
			}
		}
	}

}

func main() {
	router := gin.Default()
	router.GET("/mangas", getMangas)
	router.POST("/mangas", postMangas)
	router.GET("/mangas/:id", getMangaByID)
	router.Run("localhost:8080")
}

//TODO: after finishing - get rid of GORM completely and use raw SQL or sqlc/sqlx
//TODO: add email validation
