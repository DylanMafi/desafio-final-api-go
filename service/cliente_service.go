package service

import "desafio-final/model"

func GetAllCliente() []model.Cliente {
	var clientes []model.Cliente
	DB.Find(&clientes)
	return clientes
}

func GetClienteByID(id string) (model.Cliente, error) {
	var cliente model.Cliente
	result := DB.First(&cliente, id)
	return cliente, result.Error
}

func CreateCliente(cliente *model.Cliente) {
	DB.Create(cliente)
}

func CountClientes() int64 {
	var count int64
	DB.Model(&model.Cliente{}).Count(&count)
	return count
}
