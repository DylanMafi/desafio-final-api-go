package router

import (
	"desafio-final/controller" // Certifique-se de que o caminho está correto

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	// Definindo as rotas
	r.GET("/clientes", controller.GetAllClientes)
	r.GET("/clientes/:id", controller.GetClienteByID)
	r.POST("/clientes", controller.CreateCliente)
	r.GET("/clientes/contar", controller.CountClientes)

	return r
}
