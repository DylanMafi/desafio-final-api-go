package service

import (
	"desafio-final/model"
	"log"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

// Função para inicializar o banco de dados
func InitDatabase() {
	var err error
	// Abre a conexão com o banco de dados SQLite
	DB, err = gorm.Open(sqlite.Open("clientes.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Falha ao conectar ao banco de dados: %v", err) // Logando o erro detalhado
		return                                                     // Retorna após o erro
	}

	log.Println("Conexão ao banco de dados bem-sucedida")

	// Realiza a migração para criar a tabela 'clientes', se ainda não existir
	err = DB.AutoMigrate(&model.Cliente{})
	if err != nil {
		log.Fatalf("Erro ao migrar a base de dados: %v", err) // Logando o erro de migração
		return
	}

	log.Println("Tabela 'clientes' criada ou já existe")
}
