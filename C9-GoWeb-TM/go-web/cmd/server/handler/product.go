package handler

import (
	. "fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/JuanDRojasC/C9-GoWeb-TM/go-web/internal/products"
	"github.com/JuanDRojasC/C9-GoWeb-TM/go-web/pkg/web"
	"github.com/gin-gonic/gin"
)

type Request interface{}

type patchRequest struct {
	Name      string   `json:"nombre"`
	Color     string   `json:"color"`
	Price     float64  `json:"precio"`
	Stock     *float64 `json:"stock"`
	Code      string   `json:"codigo"`
	Published *bool    `json:"publicado"`
}

type fullRequest struct {
	Name      string   `json:"nombre" binding:"required"`
	Color     string   `json:"color" binding:"required"`
	Price     float64  `json:"precio" binding:"required"`
	Stock     *float64 `json:"stock" binding:"required"`
	Code      string   `json:"codigo" binding:"required"`
	Published *bool    `json:"publicado" binding:"required"`
}

type ProductHandler struct {
	service products.Service
}

// Get all data from service
func (ph *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		all, err := ph.service.GetAll()
		if !ph.AuthToken(ctx) {
			return
		}
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, all, nil))
	}
}

// Get one resource from service
func (ph *ProductHandler) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 || !ph.AuthToken(ctx) {
			return
		}
		p, err := ph.service.GetOne(id)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, p, nil))

	}
}

// Save a new resource
func (ph *ProductHandler) Save() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *fullRequest
		if !ph.JSONToStruct(&req, ctx) || !ph.AuthToken(ctx) {
			return
		}
		p, err := ph.service.Save(req.Name, req.Color, req.Price, *req.Stock, req.Code, *req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, p, nil))
	}
}

// Update a resource
func (ph *ProductHandler) Update() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *fullRequest
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 || !ph.JSONToStruct(&req, ctx) || !ph.AuthToken(ctx) {
			return
		}
		p, err := ph.service.Update(id, req.Name, req.Color, req.Price, *req.Stock, req.Code, *req.Published)
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, p, nil))
	}
}

// Patch a resource
func (ph *ProductHandler) Patch() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		var req *patchRequest
		if id == 0 || !ph.JSONToStruct(&req, ctx) || !ph.AuthToken(ctx) {
			return
		}

		var errs []error
		var p products.Product
		trySetProperty := func(pUpdated products.Product, err error) {
			if err != nil {
				errs = append(errs, err)
			} else {
				p = pUpdated
			}
		}

		if req.Name != "" {
			trySetProperty(ph.service.UpdateName(id, req.Name))
		}
		if req.Price != 0 {
			trySetProperty(ph.service.UpdatePrice(id, req.Price))
		}

		if len(errs) != 0 {
			ctx.JSON(http.StatusConflict, web.NewReponse(http.StatusConflict, p, errs[0]))
		} else if p == (products.Product{}) {
			ctx.JSON(http.StatusOK, "nothing modified")
		} else {
			ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, p, nil))
		}
	}
}

// Delete a resoruce
func (ph *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 || !ph.AuthToken(ctx) {
			return
		}
		if err := ph.service.Delete(id); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, "Resource Deleted", nil))
	}
}

// Validate token
func (ph *ProductHandler) AuthToken(ctx *gin.Context) bool {
	token := ctx.Request.Header.Get("token")
	if token != os.Getenv("TOKEN") {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewReponse(http.StatusBadRequest, nil, Errorf("invalid authentication")))
		return false
	}
	return true
}

// Use ShouldBindJSON, abstract the conversion returning a bool and setting in context the error if can not do it
func (ph *ProductHandler) JSONToStruct(req Request, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewReponse(http.StatusBadRequest, nil, err))
		return false
	}
	return true
}

// Convert string receive for param url in an integer that return, 0 if can not do it
func (ph *ProductHandler) ValidateIntParam(id string, ctx *gin.Context) int {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		ctx.JSON(
			http.StatusBadRequest,
			web.NewReponse(http.StatusBadRequest, nil, err))
		return 0
	}
	return idInt
}

// Retrun ProductHandler Interface
func NewProductHandler(p products.Service) *ProductHandler {
	return &ProductHandler{p}
}
