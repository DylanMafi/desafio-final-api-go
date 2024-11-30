package model

import (
	"gorm.io/gorm"
)

// Cliente representa o modelo de cliente no banco de dados.
type Cliente struct {
	gorm.Model        // Inclui campos padrão como ID, CreatedAt, UpdatedAt, DeletedAt
	Nome       string `json:"nome" gorm:"not null"`         // Campo obrigatório
	Email      string `json:"email" gorm:"unique;not null"` // Campo obrigatório e único
}
