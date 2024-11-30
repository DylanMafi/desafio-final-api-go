package controller

import (
	"desafio-final/model"
	"desafio-final/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Função para listar todos os clientes
func GetAllClientes(c *gin.Context) {
	clientes := service.GetAllClientes() // Obtém todos os clientes do serviço
	c.JSON(http.StatusOK, clientes)
}

// Função para obter um cliente específico pelo ID
func GetClienteByID(c *gin.Context) {
	id := c.Param("id")

	// Convertendo o ID de string para uint
	idUint, err := strconv.ParseUint(id, 10, 32)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "ID inválido"})
		return
	}

	cliente, err := service.GetClienteByID(uint(idUint))
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

// Função para criar um novo cliente
func CreateCliente(c *gin.Context) {
	var cliente model.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreateCliente(&cliente)
	c.JSON(http.StatusCreated, cliente)
}

// Função para contar o número total de clientes
func CountClientes(c *gin.Context) {
	count := service.CountClientes()
	c.JSON(http.StatusOK, gin.H{"total": count})
}
