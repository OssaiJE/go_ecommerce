package controllers

import (
	"context"
	"go_ecommerce/services"
	"go_ecommerce/utilities"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, _ := c.Get("user_id")
	user_id, _ := val.(primitive.ObjectID) // Type assertion
	user, _ := services.FindUserById(ctx, user_id)
	defer cancel()

	c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "User retrieved!", Data: map[string]interface{}{"user": user}})
}

func UpdateUser(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, _ := c.Get("user_id")
	user_id, _ := val.(primitive.ObjectID) // Type assertion
    // TODO
}

func UpdateProfilePhoto(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	val, _ := c.Get("user_id")
	user_id, _ := val.(primitive.ObjectID) // Type assertion
	file, _ := c.FormFile("image")
	dst := "uploads/"
	os.MkdirAll(dst, os.ModePerm)
	// Check file extension to ensure it's an image
	allowedExtensions := map[string]bool{".jpg": true, ".jpeg": true, ".png": true}
	extension := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedExtensions[extension] {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Only image files are allowed!"})
		return
	}
	// Get the current time including nanoseconds
	currentTime := time.Now().Local().Format("20060102150405.000000000")
	filename := strings.ReplaceAll(strings.ToLower(file.Filename), " ", "_")
	newFilename := user_id.Hex() + "_" + currentTime + "_" + filename

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
	update.Image = dst + newFilename
	update.Updated_At = time.Now().Local()
	user, err := services.UpdateUserById(ctx, user_id, update)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Failed to upload profile picture."})
		return
	}
	defer cancel()

	c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "User image updated!", Data: map[string]interface{}{"user": user}})
}
