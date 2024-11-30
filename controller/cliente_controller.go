package controller

import (
	"desafio-final/model"
	"desafio-final/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllClientes(c *gin.Context) {
	clientes := service.GetAllCliente()
	c.JSON(http.StatusOK, clientes)
}

func GetClienteByID(c *gin.Context) {
	id := c.Param("id")
	cliente, err := service.GetClienteByID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Cliente não encontrado"})
		return
	}
	c.JSON(http.StatusOK, cliente)
}

func CreateCliente(c *gin.Context) {
	var cliente model.Cliente
	if err := c.ShouldBindJSON(&cliente); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	service.CreateCliente(&cliente)
	c.JSON(http.StatusCreated, cliente)
}

func CountClientes(c *gin.Context) {
	count := service.CountClientes()
	c.JSON(http.StatusOK, gin.H{"total": count})
}
