package models

type Hubber struct {
	ID   string `gorm:"primaryKey"`
	Code string
	Name string
}
