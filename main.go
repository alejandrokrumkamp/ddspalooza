package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

var serialID int
var events []Event

type Event struct {
	ID              int    `json:"id"`
	Name            string `json:"name"`
	City            string `json:"city"`
	Venue           string `json:"venue"`
	PublicationDate string `json:"publicationDate"`
	EventDate       string `json:"eventDate"`
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

	r.PATCH("/events/:eventId", func(c *gin.Context) {

		eventID, _ := strconv.Atoi(c.Param("eventId"))
		for i := range events {
			if events[i].ID == eventID {
				c.BindJSON(&events[i])
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
