package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/superjantung/pet-adoption-api/models"
)

type CreateAdopterInput struct {
	Name              string `json:"name" binding:"required"`
	Email             string `json:"email" binding:"required,email"`
	Phone             string `json:"phone" binding:"required"`
	Address           string `json:"address" binding:"required"`
	ApplicationStatus string `json:"application_status" binding:"required"`
}

type UpdateAdopterInput struct {
	Name              string `json:"name"`
	Email             string `json:"email" binding:"email"`
	Phone             string `json:"phone"`
	Address           string `json:"address"`
	ApplicationStatus string `json:"application_status"`
}

func FindAdopters(c *gin.Context) {
	var adopters []models.Adopter

	if err := models.DB.Find(&adopters).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve adopters"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": adopters})
}

func FindAdopter(c *gin.Context) {
	var adopter models.Adopter

	if err := models.DB.First(&adopter, c.Param("id")).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Adopter not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch adopter"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": adopter})
}

func CreateAdopter(c *gin.Context) {
	var input CreateAdopterInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	adopter := models.Adopter{
		Name:              input.Name,
		Email:             input.Email,
		Phone:             input.Phone,
		Address:           input.Address,
		ApplicationStatus: input.ApplicationStatus,
	}

	if err := models.DB.Create(&adopter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    adopter,
		"message": "Adopter created successfully",
	})
}

func UpdateAdopter(c *gin.Context) {
	var adopter models.Adopter

	if err := models.DB.Find(&adopter, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdateAdopterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(&adopter).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update adopter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": adopter})
}

func DeleteAdopter(c *gin.Context) {
	var adopter models.Adopter

	if err := models.DB.Where("id = ?", c.Param("id")).First(&adopter).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := models.DB.Delete(&adopter).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete adopter"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
