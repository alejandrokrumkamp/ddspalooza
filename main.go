package main

import (
	"github.com/gin-gonic/gin"
)

var serialID uint32
var events []Event

type Event struct {
	ID              uint32 `json:"id"`
	Name            string `json:"name" binding:"required"`
	City            string `json:"city" binding:"required"`
	Venue           string `json:"venue" binding:"required"`
	PublicationDate string `json:"publicationDate" binding:"required"`
	EventDate       string `json:"eventDate" binding:"required"`
}

func main() {
	r := gin.Default()

	r.GET("/events", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"events": events})
	})

	r.POST("/events", func(c *gin.Context) {
		var event Event
		event.ID = serialID
		serialID++
		c.BindJSON(&event)
		events = append(events, event)
		c.JSON(200, event)
	})

	r.Run()
}
