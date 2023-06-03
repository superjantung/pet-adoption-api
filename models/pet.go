// models/pet.go

package models

import (
	"time"
)

type Pet struct {
	ID                   uint      `gorm:"primaryKey" json:"id"`
	Name                 string    `json:"name"`
	Species              string    `json:"species"`
	Breed                string    `json:"breed"`
	Age                  int       `json:"age"`
	Gender               string    `json:"gender"`
	Size                 string    `json:"size"`
	Color                string    `json:"color"`
	Weight               float64   `json:"weight"`
	Description          string    `json:"description"`
	Availability         bool      `json:"availability"`
	Location             string    `json:"location"`
	Vaccinated           bool      `json:"vaccinated"`
	MedicalHistory       string    `json:"medical_history"`
	SpecialNeeds         string    `json:"special_needs"`
	AdoptionStatus       string    `json:"adoption_status"`
	AdoptionFee          float64   `json:"adoption_fee"`
	AdoptionRequirements string    `json:"adoption_requirements"`
	CreatedAt            time.Time `json:"created_at"`
	UpdatedAt            time.Time `json:"updated_at"`
}
