package controllers

import (
	"context"
	"go_ecommerce/services"
	"go_ecommerce/utilities"
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *gin.Context) {

}

func UpdateUser(c *gin.Context) {

}

func UpdateProfilePhoto(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, _ := c.Get("user_id")
	user_id, _ := val.(primitive.ObjectID) // Type assertion
	file, _ := c.FormFile("image")
	// user_id.Hex()
	log.Println(file.Filename)
	dst := "uploads/"
	os.MkdirAll(dst, os.ModePerm)
	// Get the current time including nanoseconds
	currentTime := time.Now().Local().Format("20060102150405.000000000") // Format: YYYYMMDDHHMMSS.000000000
	filename := strings.ReplaceAll(strings.ToLower(file.Filename), " ", "_")
	newFilename := currentTime + "_" + filename
	// Upload the file to the specific destination.
	if err := c.SaveUploadedFile(file, dst+newFilename); err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Failed to upload profile picture."})
		return
	}
	type Update struct {
		Image      string
		Updated_At time.Time
	}
	var update Update
	update.Image = dst+user_id.Hex()+"_"+newFilename
	update.Updated_At = time.Now().Local()
	user, err := services.UpdateUser(ctx, user_id, update)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Failed to upload profile picture."})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "User image updated!", Data: map[string]interface{}{"user": user}})
}
