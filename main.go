package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/killerasus/gorest/godriver"
)

var database godriver.Database

func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Drivers)
}

func getDriverById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
	}
	for _, d := range database.Drivers {
		if d.ID == id {
			c.IndentedJSON(http.StatusOK, d)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Driver not found"})
}

func createDriver(c *gin.Context) {
	var driver godriver.Driver
	if err := c.BindJSON(&driver); err != nil {
		return
	}
	database.Drivers = append(database.Drivers, driver)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Driver created"})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "CAR", "message": "Hello CAR World!"})
	})
	r.GET("/drivers", getDrivers)
	r.GET("/drivers/:id", getDriverById)
	r.POST("/drivers", createDriver)

	r.Run("localhost:8080")
}
