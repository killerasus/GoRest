package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

type Driver struct {
	ID        int       `json:"id"`
	Name      string    `json:"name"`
	Birthdate time.Time `json:"birthdate"`
}

var drivers []Driver //@TODO: Use database instead

func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, drivers)
}

func getDriverById(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
	}
	for _, d := range drivers {
		if d.ID == id {
			c.IndentedJSON(http.StatusOK, d)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Driver not found"})
}

func createDriver(c *gin.Context) {
	var driver Driver
	if err := c.BindJSON(&driver); err != nil {
		return
	}
	drivers = append(drivers, driver)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "Driver created"})
}

func main() {
	r := gin.Default()
	r.LoadHTMLGlob("templates/*")
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "home.html", gin.H{"title": "CAR", "message": "Hello CAR World!"})
	})
	r.GET("/driver", getDrivers)
	r.GET("/driver/:id", getDriverById)
	r.POST("/driver", createDriver)

	r.Run("localhost:8080")
}
