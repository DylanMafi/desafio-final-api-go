package model

type Cliente struct {
	ID    uint   `json:"id" gorm:"primaryKey"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}
