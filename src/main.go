package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// @title Glowing Umbrella API
// @version 1.0.0
// @description CRM deals read API
// @host localhost:8080
// @BasePath /
func main() {
	initDB()

	r := gin.Default()
	r.GET("/deals", getDeals)
	r.Run()
}

// @Summary List deals
// @Description Returns all deals with fully nested graph
// @Tags deals
// @Produce json
// @Success 200 {array} Deal
// @Failure 500 {object} map[string]string
// @Router /deals [get]
func getDeals(c *gin.Context) {
	deals, err := GetDeals()
	if err != nil {
		log.Printf("GetDeals: %v", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}
	c.JSON(http.StatusOK, deals)
}
