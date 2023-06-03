// models/adoption.go

package models

import "time"

type Adoption struct {
	ID           uint      `gorm:"primaryKey" json:"id"`
	PetID        uint      `json:"pet_id"`
	AdopterID    uint      `json:"adopter_id"`
	AdoptionDate time.Time `json:"adoption_date"`
}
