package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/killerasus/gorest/godriver"
)

var database godriver.Database

func getDrivers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Drivers)
}

func getDriverById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	d, err := database.GetDriver(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, d)
}

func createDriver(c *gin.Context) {
	var driver godriver.Driver
	if err := c.BindJSON(&driver); err != nil {
		return
	}
	driver = database.CreateDriver(driver)
	c.IndentedJSON(http.StatusCreated, driver)
}

func updateDriver(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var driver godriver.Driver
	if err := c.BindJSON(&driver); err != nil {
		return
	}

	d, err := database.UpdateDriver(id, driver)
	if err == nil {
		c.IndentedJSON(http.StatusCreated, d)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
}

func patchDriver(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var patch godriver.DriverPatch
	if err := c.BindJSON(&patch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err = database.PatchDriver(id, patch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	d, _ := database.GetDriver(id)
	c.IndentedJSON(http.StatusAccepted, d)
}

func removeDriver(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	if _, err := database.RemoveDriver(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Driver " + id.String() + " removed"})
	}

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
	r.PUT("/drivers/:id", updateDriver)
	r.PATCH("/drivers/:id", patchDriver)
	r.DELETE("/drivers/:id", removeDriver)

	r.Run("localhost:8080")
}
