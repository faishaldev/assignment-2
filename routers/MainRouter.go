package routers

import "github.com/gin-gonic/gin"

func StartServer() *gin.Engine {
	router := gin.Default()

	OrdersRouter(*router)

	return router
}
