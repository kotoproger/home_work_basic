package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (p *Product) RegisterRoutes(router *gin.Engine) {
	apiProduct := router.Group("/api/v1/product")
	{
		apiProduct.GET("/list", func(ctx *gin.Context) {
			list, err := p.GetList(p.app.Ctx)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, list)
		})
	}
}
