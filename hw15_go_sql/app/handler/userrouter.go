package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (u *User) RegisterRoutes(router *gin.Engine) {
	apiUser := router.Group("/api/v1/user")
	{
		apiUser.POST("/register", func(ctx *gin.Context) {
			request := struct {
				Name     string `json:"name"`
				Email    string `json:"email"`
				Password string `json:"passwd"`
			}{}
			ctx.ShouldBindJSON(&request)
			user, err := u.Register(u.app.Ctx, UserDto{
				Name:     request.Name,
				Email:    request.Email,
				password: request.Password,
			})
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}
			ctx.JSON(http.StatusOK, user)
		})

		apiUser.PATCH("/name", func(ctx *gin.Context) {
			value, exists := ctx.Get("userID")
			if !exists {
				ctx.JSON(http.StatusForbidden, struct{}{})
				return
			}

			stringUserIDValue, ok := value.(string)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, struct{}{})
			}

			request := struct {
				Name string `json:"name"`
			}{}
			ctx.ShouldBindJSON(&request)
			err := u.UpdateName(
				stringUserIDValue,
				request.Name,
			)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}

			ctx.JSON(http.StatusOK, struct{}{})
		})

		apiUser.DELETE("", func(ctx *gin.Context) {
			value, exists := ctx.Get("userID")
			if !exists {
				ctx.JSON(http.StatusForbidden, struct{}{})
				return
			}

			stringUserIDValue, ok := value.(string)
			if !ok {
				ctx.JSON(http.StatusInternalServerError, struct{}{})
			}

			err := u.DeleteUser(stringUserIDValue)
			if err != nil {
				ctx.JSON(http.StatusInternalServerError, err)
				return
			}

			ctx.JSON(http.StatusOK, struct{}{})
		})
	}
}
