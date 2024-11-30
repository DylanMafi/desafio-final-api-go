package service

import "desafio-final/model"

// Função para obter todos os clientes
func GetAllClientes() []model.Cliente {
	var clientes []model.Cliente
	// Usando Find do GORM para buscar todos os registros
	DB.Find(&clientes)
	return clientes
}

// Função para obter um cliente específico pelo ID
func GetClienteByID(id uint) (model.Cliente, error) { // Usando uint para ID, que é o tipo padrão do GORM
	var cliente model.Cliente
	// Usando First para buscar o primeiro cliente com o ID fornecido
	result := DB.First(&cliente, id)
	return cliente, result.Error
}

// Função para criar um novo cliente
func CreateCliente(cliente *model.Cliente) error {
	// Usando Create para adicionar o cliente ao banco de dados
	result := DB.Create(cliente)
	return result.Error // Retorna o erro caso ocorra
}

// Função para contar o número total de clientes
func CountClientes() int64 {
	var count int64
	// Usando o método Count para contar os registros na tabela de clientes
	DB.Model(&model.Cliente{}).Count(&count)
	return count
}
