package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (o *Order) RegisterRoutes(router *gin.Engine) {
	apiOrder := router.Group("/api/v1/order")
	{
		apiOrder.GET("/:orderID", func(ctx *gin.Context) {
			value, exists := ctx.Get("userID")
			if !exists {
				ctx.JSON(http.StatusForbidden, struct{}{})
				return
			}
			stringUserIDValue, ok := value.(string)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, struct{}{})
			}
			stringOrderIDValue := ctx.Param("orderID")
			if stringOrderIDValue == "" {
				ctx.JSON(http.StatusBadRequest, struct{}{})
				return
			}
			order, err := o.GetByID(o.app.Ctx, stringOrderIDValue, stringUserIDValue)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, order)
		})
	}
}
