package routers

import (
	. "assignment-2/controllers"

	"github.com/gin-gonic/gin"
)

func OrderRouter(router gin.Engine) *gin.Engine {
	router.POST("/orders", PostOrder)
	router.GET("/orders", GetOrders)
	router.PUT("/orders/:orderId", PutOrderById)

	return &router
}
