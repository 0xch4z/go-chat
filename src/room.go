package main

import (
	"fmt"
	"io"

	"github.com/dustin/go-broadcast"
	"github.com/gin-gonic/gin"
)

var rooms = make(map[string]Room)

// Listener Type
type Listener chan interface{}

// Room type
type Room struct {
	broadcaster broadcast.Broadcaster
	messages    []string
}

func getRoom(roomID string) Room {
	room, ok := rooms[roomID]
	if !ok {
		broadcaster := broadcast.NewBroadcaster(10)
		messages := []string{}
		room = Room{broadcaster, messages}
		rooms[roomID] = room
	}
	return room
}

func openListener(roomID string) Listener {
	listener := make(Listener)
	getRoom(roomID).broadcaster.Register(listener)
	return listener
}

func closeListener(roomID string, listener Listener) {
	getRoom(roomID).broadcaster.Unregister(listener)
	close(listener)
}

func deleteBroadcast(roomID string) {
	room, ok := rooms[roomID]
	fmt.Printf("DELETING ROOM %s!", room)
	if ok {
		room.broadcaster.Close()
		delete(rooms, roomID)
	}
}

/// Routes

func roomGET(c *gin.Context) {
	roomID := c.Param("room_id")
	userID := randomKey()

	room := getRoom(roomID)

	fmt.Printf("there are %d messages", len(room.messages))

	render(c, gin.H{
		"user_id":          userID,
		"room_id":          roomID,
		"initial_messages": room.messages,
	}, "index.tmpl")
}

func roomPOST(c *gin.Context) {
	roomID := c.Param("room_id")
	userID := c.PostForm("user_id")
	message := c.PostForm("message")

	sum := fmt.Sprintf("%s: %s", userID, message)

	room := getRoom(roomID)
	room.broadcaster.Submit(sum)
	room.messages = append(room.messages, sum)
	rooms[roomID] = room

	c.JSON(201, gin.H{
		"success": true,
		"message": message,
	})
}

func roomDELETE(c *gin.Context) {
	roomID := c.Param("room_id")
	deleteBroadcast(roomID)
}

func stream(c *gin.Context) {
	roomID := c.Param("room_id")
	listener := openListener(roomID)
	defer closeListener(roomID, listener)

	c.Stream(func(w io.Writer) bool {
		c.SSEvent("message", <-listener)
		return true
	})
}
