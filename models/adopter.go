// models/adopter.go

package models

type Adopter struct {
	ID                uint       `gorm:"primaryKey" json:"id"`
	Name              string     `json:"name"`
	Email             string     `json:"email"`
	Phone             string     `json:"phone"`
	Address           string     `json:"address"`
	ApplicationStatus string     `json:"application_status"`
	Adoptions         []Adoption `gorm:"foreignkey:AdopterID" json:"adoptions"`
}
