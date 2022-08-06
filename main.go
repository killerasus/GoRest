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

func getPassengers(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, database.Passengers)
}

func getPassengerById(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	p, err := database.GetPassenger(id)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusOK, p)
}

func createPassenger(c *gin.Context) {
	var passenger godriver.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		return
	}
	passenger = database.CreatePassenger(passenger)
	c.IndentedJSON(http.StatusCreated, passenger)
}

func updatePassenger(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var passenger godriver.Passenger
	if err := c.BindJSON(&passenger); err != nil {
		return
	}

	p, err := database.UpdatePassenger(id, passenger)
	if err == nil {
		c.IndentedJSON(http.StatusCreated, p)
	} else {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	}
}

func patchPassenger(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	var patch godriver.PassengerPatch
	if err := c.BindJSON(&patch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	if err = database.PatchPassenger(id, patch); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	p, _ := database.GetPassenger(id)
	c.IndentedJSON(http.StatusAccepted, p)
}

func removePassenger(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid id"})
		return
	}

	if _, err := database.RemovePassenger(id); err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
	} else {
		c.IndentedJSON(http.StatusAccepted, gin.H{"message": "Passenger " + id.String() + " removed"})
	}
}

func createTravelRequest(c *gin.Context) {
	var travel godriver.TravelRequestInput
	if err := c.BindJSON(&travel); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": err})
		return
	}

	switch {
	case travel.Passenger == nil:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Passenger field missing"})
		return
	case travel.Origin == nil:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Origin field missing"})
		return
	case travel.Destination == nil:
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Destination field missing"})
		return
	}

	created, err := database.SaveTravelRequest(travel)
	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": err})
		return
	}

	c.IndentedJSON(http.StatusAccepted, created)
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

	r.GET("/passengers", getPassengers)
	r.GET("/passengers/:id", getPassengerById)
	r.POST("/passengers", createPassenger)
	r.PUT("/passengers/:id", updatePassenger)
	r.PATCH("/passengers/:id", patchPassenger)
	r.DELETE("/passengers/:id", removePassenger)

	r.POST("/travelRequests", createTravelRequest)

	r.Run("localhost:8080")
}
