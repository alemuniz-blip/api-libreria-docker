package api

import (
	"database/sql"
	"fmt"
	"net/http"

	"rest/dto"

	"github.com/gin-gonic/gin"
)

// CREATE PRODUCT
func (server *Server) createProduct(ctx *gin.Context) {

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int32   `json:"stock"`
		CategoryID  int32   `json:"category_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.CreateProduct(ctx, dto.CreateProductParams{
		Name: req.Name,

		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},

		Price: sql.NullString{
			String: fmt.Sprintf("%.2f", req.Price),
			Valid:  true,
		},

		Stock: sql.NullInt32{
			Int32: req.Stock,
			Valid: true,
		},
		CategoryID: req.CategoryID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Producto creado correctamente",
	})
}

// UPDATE PRODUCT
func (server *Server) updateProduct(ctx *gin.Context) {

	idParam := ctx.Param("id")

	var id int32
	fmt.Sscan(idParam, &id)

	var req struct {
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		Stock       int32   `json:"stock"`
		CategoryID  int32   `json:"category_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.UpdateProduct(ctx, dto.UpdateProductParams{
		Name: req.Name,

		Description: sql.NullString{
			String: req.Description,
			Valid:  true,
		},

		Price: sql.NullString{
			String: fmt.Sprintf("%.2f", req.Price),
			Valid:  true,
		},

		Stock: sql.NullInt32{
			Int32: req.Stock,
			Valid: true,
		},

		CategoryID: req.CategoryID,
		ID:         id,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Producto actualizado correctamente",
	})
}

// GET ALL PRODUCTS
func (server *Server) getAllProducts(ctx *gin.Context) {

	products, err := server.dbtx.GetAllProducts(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, p := range products {
		response = append(response, gin.H{
			"id":          p.ID,
			"name":        p.Name,
			"description": p.Description.String,
			"price":       p.Price.String,
			"stock":       p.Stock.Int32,
			"category":    p.Category,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// GET PRODUCTS BY CATEGORY
func (server *Server) getProductsByCategory(ctx *gin.Context) {

	categoryParam := ctx.Param("id")

	var categoryID int32
	fmt.Sscan(categoryParam, &categoryID)

	products, err := server.dbtx.GetProductsByCategory(ctx, categoryID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, p := range products {
		response = append(response, gin.H{
			"id":          p.ID,
			"name":        p.Name,
			"description": p.Description.String,
			"price":       p.Price.String,
			"stock":       p.Stock.Int32,
			"category":    p.Category,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// DELETE PRODUCT
func (server *Server) deleteProduct(ctx *gin.Context) {

	idParam := ctx.Query("id")

	var id int32
	fmt.Sscan(idParam, &id)

	err := server.dbtx.DeleteProduct(ctx, id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "Producto eliminado correctamente",
	})
}
