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
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product models.Product
	var user_id primitive.ObjectID
	val, _ := c.Get("user_id")
	if val != nil {
		// Type assertion
		user_id, _ = val.(primitive.ObjectID)
	}

	if err := c.BindJSON(&product); err != nil {
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Bad request!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	product.User_ID = user_id
	product.Created_At = time.Now().Local()
	product.Updated_At = time.Now().Local()
	product_id, err := services.CreateProduct(ctx, product)
	if product_id == nil {
		log.Panic(err)
		c.JSON(http.StatusInternalServerError, utilities.UserResponse{Status: http.StatusInternalServerError, Message: "Failed to create product!", Data: map[string]interface{}{"data": err.Error()}})
		return
	}
	defer cancel()
	c.JSON(http.StatusCreated, utilities.UserResponse{Status: http.StatusCreated, Message: "Product created!", Data: map[string]interface{}{"product_id": product_id}})
}

func GetOneProduct(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var product *models.Product
	id := c.Param("id")
	product_id, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "Invalid product ID!"})
		return
	}
	product, err = services.FindProductById(ctx, product_id)
	defer cancel()

	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "No product found!"})
		return
	}

	c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "Product retrieved!", Data: map[string]interface{}{"product": product}})

}

func GetAllProducts(c *gin.Context) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var products []*models.Product

	products, err := services.FindAllProducts(ctx)
    if err != nil || len(products) < 1 {
        log.Println(err, len(products))
		c.JSON(http.StatusBadRequest, utilities.UserResponse{Status: http.StatusBadRequest, Message: "No product found!"})
		return
    }
    defer cancel()
    c.JSON(http.StatusOK, utilities.UserResponse{Status: http.StatusOK, Message: "Products retrieved!", Data: map[string]interface{}{"products": products}})
}

func UpdateProduct(c *gin.Context) {

}

func DeleteProduct(c *gin.Context) {

}
