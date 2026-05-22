package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	initDB()

	r := gin.Default()
	r.GET("/deals", func(c *gin.Context) {
		deals, err := GetDeals()
		if err != nil {
			log.Printf("GetDeals: %v", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}
		c.JSON(http.StatusOK, deals)
	})
	r.Run()
}
