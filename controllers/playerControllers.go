package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/zachmshort/monopoly-backend/config"
	"github.com/zachmshort/monopoly-backend/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetPlayersInRoom(c *gin.Context) {
	roomCode := c.Param("roomCode")

	roomCollection := config.DB.Collection("Room")
	playerCollection := config.DB.Collection("Player")

	var room models.Room
	err := roomCollection.FindOne(c, bson.M{"roomCode": roomCode}).Decode(&room)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Room not found"})
		return
	}

	query := bson.M{"roomId": room.ID}

	var players []models.Player
	cursor, err := playerCollection.Find(c, query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get players"})
		return
	}

	if err = cursor.All(c, &players); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode players"})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"players": players,
		"room":    room,
	})
}

func GetPlayerByDevice(c *gin.Context) {
	deviceId := c.Param("deviceId")

	collection := config.DB.Collection("Player")
	var player models.Player
	err := collection.FindOne(c, bson.M{"deviceId": deviceId}).Decode(&player)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	c.JSON(http.StatusOK, player)
}

func GetPlayerDetails(c *gin.Context) {
	playerId, _ := primitive.ObjectIDFromHex(c.Param("playerId"))

	playerColl := config.DB.Collection("Player")
	var player models.Player
	err := playerColl.FindOne(c, bson.M{"_id": playerId}).Decode(&player)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Player not found"})
		return
	}

	propColl := config.DB.Collection("Property")
	var properties []models.Property
	cursor, err := propColl.Find(c, bson.M{"playerId": playerId})
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to get properties"})
		return
	}

	if err = cursor.All(c, &properties); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to decode properties"})
		return
	}

	response := map[string]interface{}{
		"player":     player,
		"properties": properties,
	}

	c.JSON(http.StatusOK, response)
}
