package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"github.com/superjantung/pet-adoption-api/models"
)

type CreatePetInput struct {
	Name                 string   `json:"name" binding:"required"`
	Species              string   `json:"species" binding:"required"`
	Breed                string   `json:"breed" binding:"required"`
	Age                  int      `json:"age" binding:"required"`
	Gender               string   `json:"gender"`
	Size                 string   `json:"size"`
	Color                string   `json:"color"`
	Weight               float64  `json:"weight"`
	Description          string   `json:"description"`
	Photos               []string `json:"photos"`
	Availability         bool     `json:"availability"`
	Location             string   `json:"location"`
	Vaccinated           bool     `json:"vaccinated"`
	MedicalHistory       string   `json:"medical_history"`
	SpecialNeeds         string   `json:"special_needs"`
	AdoptionStatus       string   `json:"adoption_status"`
	AdoptionFee          float64  `json:"adoption_fee"`
	AdoptionRequirements string   `json:"adoption_requirements"`
}
type UpdatePetInput struct {
	Name                 string   `json:"name"`
	Species              string   `json:"species"`
	Breed                string   `json:"breed"`
	Age                  int      `json:"age"`
	Gender               string   `json:"gender"`
	Size                 string   `json:"size"`
	Color                string   `json:"color"`
	Weight               float64  `json:"weight"`
	Description          string   `json:"description"`
	Photos               []string `json:"photos"`
	Availability         bool     `json:"availability"`
	Location             string   `json:"location"`
	Vaccinated           bool     `json:"vaccinated"`
	MedicalHistory       string   `json:"medical_history"`
	SpecialNeeds         string   `json:"special_needs"`
	AdoptionStatus       string   `json:"adoption_status"`
	AdoptionFee          float64  `json:"adoption_fee"`
	AdoptionRequirements string   `json:"adoption_requirements"`
}

func FindPets(c *gin.Context) {
	var pets []models.Pet
	if err := models.DB.Find(&pets).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve pets"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pets})
}

func FindPet(c *gin.Context) {
	var pet models.Pet

	if err := models.DB.First(&pet, c.Param("id")).Error; err != nil {
		if gorm.IsRecordNotFoundError(err) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Pet not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch pet"})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pet})
}

func CreatePet(c *gin.Context) {
	var input CreatePetInput

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	pet := models.Pet{
		Name:                 input.Name,
		Species:              input.Species,
		Breed:                input.Breed,
		Age:                  input.Age,
		Gender:               input.Gender,
		Size:                 input.Size,
		Color:                input.Color,
		Weight:               input.Weight,
		Description:          input.Description,
		Availability:         input.Availability,
		Location:             input.Location,
		Vaccinated:           input.Vaccinated,
		MedicalHistory:       input.MedicalHistory,
		SpecialNeeds:         input.SpecialNeeds,
		AdoptionStatus:       input.AdoptionStatus,
		AdoptionFee:          input.AdoptionFee,
		AdoptionRequirements: input.AdoptionRequirements,
	}

	if err := models.DB.Create(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":    pet,
		"message": "Pet created successfully",
	})
}

func UpdatePet(c *gin.Context) {
	var pet models.Pet

	if err := models.DB.Find(&pet, c.Param("id")).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	var input UpdatePetInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if err := models.DB.Model(&pet).Updates(input).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update pet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": pet})
}

func DeletePet(c *gin.Context) {

	var pet models.Pet
	if err := models.DB.Where("id = ?", c.Param("id")).First(&pet).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	if err := models.DB.Delete(&pet).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete pet"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": true})
}
