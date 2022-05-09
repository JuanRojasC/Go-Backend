package handler

import (
	"os"

	"github.com/gin-gonic/gin"
	"github.com/ignaciofalco/dynamo-storage/internal/users"
)

type request struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

type User struct {
	service users.Service
}

func NewUser(u users.Service) *User {
	return &User{
		service: u,
	}
}

func (c *User) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		p, err := c.service.GetOne(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *User) Store() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{"error": "token inválido"})
			return
		}

		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Store(req.Firstname, req.Lastname, req.Username, req.Email)
		if err != nil {
			ctx.JSON(404, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(200, p)
	}
}

func (c *User) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		token := ctx.Request.Header.Get("token")
		tokenFromEnv := os.Getenv("TOKEN")

		if token != tokenFromEnv {
			ctx.JSON(401, gin.H{
				"error": "token inválido",
			})
			return
		}

		err := c.service.Delete(ctx.Param("id"))
		if err != nil {
			ctx.JSON(500, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.Status(204)
	}
}
