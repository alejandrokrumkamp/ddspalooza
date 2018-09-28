package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var serialID int
var events []Event

type Event struct {
	ID              int    `json:"id"`
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

	r.GET("/events/:eventId", func(c *gin.Context) {
		eventID, _ := strconv.Atoi(c.Param("eventId"))
		for i := range events {
			if events[i].ID == eventID {
				c.JSON(200, events[i])
				return
			}
		}

		c.JSON(404, gin.H{"message": "Event not found"})
	})

	r.POST("/events", func(c *gin.Context) {
		var event Event
		event.ID = serialID
		serialID++
		c.BindJSON(&event)
		events = append(events, event)
		c.JSON(200, event)
	})

	// TODO prevent already set fields to be overwritten with zero values
	r.PATCH("/events/:eventId", func(c *gin.Context) {
		var event Event
		c.BindJSON(&event)

		eventID, _ := strconv.Atoi(c.Param("eventId"))
		for i := range events {
			if events[i].ID == eventID {
				event.ID = events[i].ID
				events[i] = event
				c.JSON(200, events[i])
				return
			}
		}
		c.JSON(404, gin.H{"message": "Event not found"})
	})

	r.DELETE("/events/:eventId", func(c *gin.Context) {
		eventID, _ := strconv.Atoi(c.Param("eventId"))
		for i := range events {
			if events[i].ID == eventID {
				events[i] = events[len(events)-1]
				var event Event
				events[len(events)-1] = event
				events = events[:len(events)-1]
				c.JSON(200, gin.H{"message": "OK"})
				return
			}
		}

		c.JSON(404, gin.H{"message": "Event not found"})
	})

	r.Run(":8080")
}
