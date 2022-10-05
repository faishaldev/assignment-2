package routers

import (
	"github.com/gin-gonic/gin"
)

func OrdersRouter(router gin.Engine) *gin.Engine {
	// router.POST("/orders", orders.CreateOrder())
	// router.GET("/orders", orders.GetOrder())
	// router.PUT("/orders/:order_id", orders.UpdateOrder())
	// router.DELETE("/orders/:order_id", orders.DeleteOrder())

	return &router
}
