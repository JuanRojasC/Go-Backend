package handler

import (
	. "fmt"
	"net/http"
	"os"
	"reflect"
	"strconv"
	"strings"

	"github.com/JuanDRojasC/C12-Testing-TT/go-web/internal/products"
	"github.com/JuanDRojasC/C12-Testing-TT/go-web/pkg/web"
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

// ListProducts godoc
// @Summary List products
// @Tags Products
// @Description get all products available
// @Produce json
// @Param token header string true "token"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response "Invalid Token"
// @Router /products [get]
func (ph *ProductHandler) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		all, err := ph.service.GetAll()
		if err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusOK, web.NewReponse(http.StatusOK, all, nil))
	}
}

// GetProduct godoc
// @Summary Product by id
// @Tags Products
// @Description get product by id
// @Accept json
// @Produce json
// @Param token header string true "token" tru "id"
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "Param ID is not an integer"
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 404 {object} web.Response "Resource not found"
// @Router /products/:id [get]
func (ph *ProductHandler) GetOne() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 {
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

// SaveProduct godoc
// @Summary Save a new product
// @Tags Products
// @Description Save a new product
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param product body fullRequest true "Product to Save"
// @Success 200 {object} web.Response
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 422 {object} web.Response "Body malformed"
// @Router /products [post]
func (ph *ProductHandler) SaveProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *fullRequest
		if !ph.JSONToStruct(&req, ctx) {
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

// UpdateProduct godoc
// @Summary Update product existing
// @Tags Products
// @Description Update product existing
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Param product body fullRequest true "Product to Update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "Param ID is not an integer"
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 404 {object} web.Response "Resource not found"
// @Failure 422 {object} web.Response "Body malformed"
// @Router /products/:id [put]
func (ph *ProductHandler) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var req *fullRequest
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 || !ph.JSONToStruct(&req, ctx) {
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

// PatchProduct godoc
// @Summary Patch one or many product's properties
// @Tags Products
// @Description Patch one or many product's properties
// @Accept json
// @Produce json
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Param product body patchRequest false "Product properties to Update"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "Param ID is not an integer"
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 404 {object} web.Response "Resource not found"
// @Failure 422 {object} web.Response "Body malformed"
// @Router /products/:id [patch]
func (ph *ProductHandler) PatchProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		var req *patchRequest
		if id == 0 || !ph.JSONToStruct(&req, ctx) {
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

// Delete Product godoc
// @Summary Delete a product existing
// @Tags Products
// @Description Delete a product existing
// @Param token header string true "token"
// @Param id path int true "Product ID"
// @Success 200 {object} web.Response
// @Failure 400 {object} web.Response "Param ID is not an integer"
// @Failure 401 {object} web.Response "Invalid Token"
// @Failure 404 {object} web.Response "Resource not found"
// @Router /products/:id [delete]
func (ph *ProductHandler) Delete() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id := ph.ValidateIntParam(ctx.Param("id"), ctx)
		if id == 0 {
			return
		}
		if err := ph.service.Delete(id); err != nil {
			ctx.JSON(http.StatusNotFound, web.NewReponse(http.StatusNotFound, nil, err))
			return
		}
		ctx.JSON(http.StatusNoContent, web.NewReponse(http.StatusNoContent, "Resource Deleted", nil))
	}
}

// Validate token
func (ph *ProductHandler) AuthToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("token")
		if token != os.Getenv("TOKEN") {
			ctx.AbortWithStatusJSON(
				http.StatusUnauthorized,
				web.NewReponse(http.StatusUnauthorized, nil, Errorf("invalid authentication")))
		}
	}
}

// Use ShouldBindJSON, abstract the conversion returning a bool and setting in context the error if can not do it
func (ph *ProductHandler) JSONToStruct(req Request, ctx *gin.Context) bool {
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(
			http.StatusUnprocessableEntity,
			web.NewReponse(http.StatusUnprocessableEntity, nil, ErrorBindingJSON(req, err)))
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

// Return a error witha custom message for client
// showing what field JSON is malformed and need to be replaced
func ErrorBindingJSON(req Request, err error) error {
	var t reflect.Type

	switch req.(type) {
	case **patchRequest:
		t = reflect.TypeOf(patchRequest{})
	default:
		t = reflect.TypeOf(fullRequest{})
	}

	fieldName := strings.Split(err.Error(), "'")[3]
	f, ok := t.FieldByName(fieldName)
	if !ok {
		return err
	}
	jsonName, jsonErr := f.Tag.Lookup("json")
	if !jsonErr {
		return err
	}
	return Errorf("property '%s' is required", jsonName)
}

// Return ProductHandler Interface
func NewProductHandler(p products.Service) *ProductHandler {
	return &ProductHandler{p}
}
