package controllers

import (
	"assignment2/database"
	"assignment2/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func CreateOrder(c *gin.Context) {
	var orderInput models.OrderInput
	var order models.Order
	var items []models.Item

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db := database.GetDB()

	order.CustomerName = orderInput.CustomerName

	items = append(items, models.Item{ItemCode: orderInput.ItemCode, Description: orderInput.Description, Quantity: orderInput.Quantity})
	order.Items = items

	if err := db.Create(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"order": order})
}

func GetOrder(c *gin.Context) {

	db := database.GetDB()
	var orders []models.Order
	if err := db.Preload("Items").Find(&orders).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{"orders": orders})
}

func UpdateOrder(c *gin.Context) {
	// Get the order ID from the URL parameter
	orderId := c.Param("orderId")

	// Get the order from the database by ID
	var orderInput models.OrderInput
	var order models.Order
	var item models.Item

	if err := c.ShouldBindJSON(&orderInput); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	dbOrder := database.GetDB()
	if err := dbOrder.Preload("Items").Where("order_id = ?", orderId).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	dbItem := database.GetDB()
	if err := dbItem.Where("order_id = ?", orderId).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	item.ItemCode = orderInput.ItemCode
	item.Description = orderInput.Description
	item.Quantity = orderInput.Quantity
	order.CustomerName = orderInput.CustomerName
	order.Items[0] = item

	dbOrder.Save(&order)
	dbOrder.Save(&item)

	// Return the updated order
	c.JSON(http.StatusOK, gin.H{"order": order})
}

func DeleteOrder(c *gin.Context) {
	// Get the order ID from the URL parameter
	orderId := c.Param("orderId")

	db := database.GetDB()

	// Get the item from the database
	var item models.Item
	if err := db.Where("order_id = ?", orderId).First(&item).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Item not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Delete the item from the database
	if err := db.Delete(&item).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Get the order from the database by ID

	var order models.Order
	if err := db.Where("order_id = ?", orderId).First(&order).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			c.JSON(http.StatusNotFound, gin.H{"error": "Order not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}

	// Delete the order from the database
	if err := db.Delete(&order).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	// Return a success message
	c.JSON(http.StatusOK, gin.H{"message": "Order deleted successfully"})
}
