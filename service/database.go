package service

import (
	"desafio-final/model"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("clinetes.db"), &gorm.Config{})
	if err != nil {
		panic("Falha ao conectar ao banco de dados")
	}
	DB.AutoMigrate(&model.Cliente{})
}
