package handler

import (
	"net/http"
	"strconv"

	"github.com/JuanDRojasC/C7-GoWeb-TT/go-web/internal/products"
	"github.com/gin-gonic/gin"
)

// 2.a
type request struct {
	Name      string   `json:"nombre" binding:"required"`
	Color     string   `json:"color" binding:"required"`
	Price     float64  `json:"precio" binding:"required"`
	Stock     *float64 `json:"stock" binding:"required"`
	Code      string   `json:"codigo" binding:"required"`
	Published *bool    `json:"publicado" binding:"required"`
}

// 2.b
type Product struct {
	service products.Service
}

// 2.d

func (c *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}

		all, err := c.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, all)
	}
}

func (c *Product) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{
				"error": "token inválido",
			})
			return
		}
		idFind := ctx.Param("id")
		if idFind == "" {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": "id param empty",
			})
			return
		}
		id, err := strconv.Atoi(idFind)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.GetOne(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{
				"error": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusOK, p)

	}
}

func (c *Product) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != "123456" {
			ctx.JSON(http.StatusUnauthorized, gin.H{"error": "token inválido"})
			return
		}
		var req request
		if err := ctx.ShouldBindJSON(&req); err != nil {
			ctx.JSON(404, gin.H{
				"error": err.Error(),
			})
			return
		}
		p, err := c.service.Save(req.Name, req.Color, req.Price, *req.Stock, req.Code, *req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, p)
	}
}

// 2.c
func NewProduct(p products.Service) *Product {
	return &Product{p}
}
