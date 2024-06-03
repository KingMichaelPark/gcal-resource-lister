package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func listRoomEvents(c *gin.Context) {
	name := c.Param("name")
	day := c.DefaultQuery("day", "today")

	var events []ParsedEvent = getRoomEvents(name, day)

	body, _ := marshal(events)
	c.Data(http.StatusOK, "application/json", body)

}

func listRooms(c *gin.Context) {
	name := c.Param("name")
	day := c.DefaultQuery("day", "today")

	var events []ParsedEvent = getRoomEvents(name, day)

	body, _ := marshal(events)
	c.Data(http.StatusOK, "application/json", body)

}

func main() {
	// var ROOMS = getAllRooms()

	r := gin.Default()
	r.GET("/rooms/:calendarId", func(c *gin.Context) {
		listRoomEvents(c)
	})
	r.GET("/rooms", func(c *gin.Context) {
		listRooms(c)
	})
	r.Run()
}
