package api

import (
	"database/sql"
	"fmt"
	"net/http"
	"time"

	"rest/dto"

	"github.com/gin-gonic/gin"
)

// CREATE COMPRA
func (server *Server) createCompra(ctx *gin.Context) {

	var req struct {
		UsuarioID int32   `json:"usuario_id"`
		Total     float64 `json:"total"`
		Estado    string  `json:"estado"`
	}
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.dbtx.CreateCompra(ctx, dto.CreateCompraParams{

		UsuarioID: sql.NullInt32{
			Int32: req.UsuarioID,
			Valid: true,
		},

		Total: sql.NullString{
			String: fmt.Sprintf("%.2f", req.Total),
			Valid:  true,
		},

		Estado: sql.NullString{
			String: req.Estado,
			Valid:  true,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastID, _ := result.LastInsertId()

	ctx.JSON(http.StatusOK, gin.H{
		"id":         lastID,
		"usuario_id": req.UsuarioID,
		"fecha":      time.Now(),
		"total":      req.Total,
		"estado":     req.Estado,
		"message":    "Compra creada 🧾",
	})
}

// cREATE DETALLE COMPRA
func (server *Server) createDetalleCompra(ctx *gin.Context) {

	var req struct {
		CompraID   int32 `json:"compra_id"`
		ProductoID int32 `json:"producto_id"`
		Cantidad   int32 `json:"cantidad"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	product, err := server.dbtx.GetProductById(ctx, req.ProductoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var precio float64
	fmt.Sscan(product.Price.String, &precio)

	subtotal := precio * float64(req.Cantidad)

	err = server.dbtx.CreateDetalleCompra(ctx, dto.CreateDetalleCompraParams{
		CompraID:   req.CompraID,
		ProductoID: req.ProductoID,
		Cantidad:   req.Cantidad,

		PrecioUnitario: sql.NullString{
			String: fmt.Sprintf("%.2f", precio),
			Valid:  true,
		},

		Subtotal: sql.NullString{
			String: fmt.Sprintf("%.2f", subtotal),
			Valid:  true,
		},
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	err = server.dbtx.UpdateCompraTotal(ctx, dto.UpdateCompraTotalParams{
		Total: sql.NullString{
			String: fmt.Sprintf("%.2f", subtotal),
			Valid:  true,
		},
		ID: req.CompraID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"compra_id":       req.CompraID,
		"producto_id":     req.ProductoID,
		"cantidad":        req.Cantidad,
		"precio_unitario": precio,
		"subtotal":        subtotal,
		"message":         "Detalle compra creado 🧾",
	})
}

// GET ALL COMPRAS
func (server *Server) getAllCompras(ctx *gin.Context) {

	compras, err := server.dbtx.GetAllCompras(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, compras)
}

// GET DETALLE COMPRA
func (server *Server) getDetalleCompra(ctx *gin.Context) {

	idParam := ctx.Param("id")

	var compraID int32
	fmt.Sscan(idParam, &compraID)

	detalles, err := server.dbtx.GetDetalleCompra(ctx, compraID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, detalles)
}

// GET COMPRAS BY USER
func (server *Server) getComprasByUser(ctx *gin.Context) {

	idParam := ctx.Param("id")

	var userID int32
	fmt.Sscan(idParam, &userID)

	compras, err := server.dbtx.GetComprasByUser(ctx, sql.NullInt32{
		Int32: userID,
		Valid: true,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, compras)
}
