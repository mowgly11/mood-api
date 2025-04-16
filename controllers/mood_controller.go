package controllers

import (
	"log"
	"net/http"
	//"fmt"
	"project/database"
	"project/models"
	"github.com/gin-gonic/gin"
)

func GetMood(c *gin.Context) {
	id := c.Query("id")
	var data models.Mood

	row, err := database.DB.Query("SELECT * FROM UserMood WHERE userId = ?", id)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	defer row.Close()

	row.Next()
	if err := row.Scan(&data.ID, &data.Mood, &data.LastUpdated); err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"mood": data})
}

func UpdateMood(c *gin.Context) {
	var data models.UpdatedMood

	if err := c.BindJSON(&data); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Invalid Content"})
		return
	}

	_, err := database.DB.Exec("UPDATE UserMood SET mood = ?, lastUpdated = ? WHERE userId = ?", data.NewMood, 0, data.ID)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Database error"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "successfully updated mood"})
}