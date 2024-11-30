package main

import (
	"desafio-final/router"  // Importando o pacote de roteamento
	"desafio-final/service" // Importando o service para inicializar o banco de dados
	"log"                   // Para logar erros e informações

	_ "github.com/mattn/go-sqlite3" // Importando o driver SQLite
)

func main() {
	// Inicialize o banco de dados antes de usar
	service.InitDatabase() // A função InitDatabase não retorna erro, então não precisa de verificação de erro

	log.Println("Banco de dados inicializado com sucesso")

	// Iniciar o servidor Gin
	r := router.SetupRouter() // Usando a função SetupRouter do pacote router
	log.Println("Iniciando o servidor na porta 8080...")
	if err := r.Run(":8080"); err != nil { // Inicia o servidor e verifica erros
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
