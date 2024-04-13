package controllers

import (
	"context"
	"go_ecommerce/models"
	"go_ecommerce/services"
	"go_ecommerce/utilities"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

var validate = validator.New()

func SignUp(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user models.User
	//validate the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Bad request!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	//use the validator library to validate required fields
	if validationErr := validate.Struct(&user); validationErr != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Validation error!", Data: map[string]interface{}{"data": validationErr.Error()}})
		return
	}
	count, err := services.UserExistByEmail(ctx, user.Email)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusInternalServerError, Message: "Server error!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "User already exist!"})
		return
	}
	count, err = services.UserExistByPhoneNumber(ctx, *user.Phone_Number)
	if err != nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusInternalServerError, Message: "Server error!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	if count > 0 {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Phone number is already in use by a user!"})
		return
	}

	user.Password, err = utilities.HashPassword(user.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Unable to process request!"})
		return
	}

	user.ID = primitive.NewObjectID()
	token, _ := services.GenerateToken(user.ID, user.First_Name, user.Last_Name, user.Email)

	user.User_Cart = make([]models.ProductUser, 0)
	user.Address_Details = make([]models.Address, 0)
	user.Order_Status = make([]models.Order, 0)
	user.Created_At = time.Now().Local()
	user.Updated_At = time.Now().Local()

	user_id, err := services.CreateUser(ctx, user)
	if user_id == nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusInternalServerError, Message: "Failed to create user!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	defer cancel()
	c.JSON(http.StatusCreated, utilities.UserResponse{Status: http.StatusCreated, Message: "Sign up successful!", Data: map[string]interface{}{"token": token, "user_id": user_id}})

}

func SignIn(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var user *models.User

	//validate the request body
	if err := c.BindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Bad request!"})
		return
	}

	found_user, err := services.FindUserByEmail(ctx, user.Email)
	defer cancel()

	if err != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Invalid login credentials!"})
		return
	}

	isValid := utilities.VerifyPassword(found_user.Password, user.Password)
	defer cancel()

	if !isValid {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Invalid login credentials!"})
		return
	}

	token, _ := services.GenerateToken(found_user.ID, found_user.First_Name, found_user.Last_Name, found_user.Email)
	defer cancel()

	c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "Sign in successful!", Data: map[string]interface{}{"token": token, "user": found_user}})
}
