package router

import (
	"desafio-final/controller"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/clientes", controller.GetAllClientes)
	r.GET("/clientes/:id", controller.GetClienteByID)
	r.POST("/clientes", controller.CreateCliente)
	r.GET("/clientes/contar", controller.CountClientes)
	return r
}
