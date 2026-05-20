package api

import (
	"fmt"
	"net/http"
	"rest/dto"
	"strconv"

	"github.com/gin-gonic/gin"
)

// CREATE CARRITO
func (server *Server) createCarrito(ctx *gin.Context) {

	var req struct {
		UsuarioID int32 `json:"usuario_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.CreateCarrito(ctx, req.UsuarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "carrito creado 🛒",
	})
}

// GET ALL CARRITOS
func (server *Server) getAllCarritos(ctx *gin.Context) {

	carritos, err := server.dbtx.GetAllCarritos(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, c := range carritos {
		response = append(response, gin.H{
			"carrito_id": c.CarritoID,
			"usuario_id": c.UsuarioID,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// GET CARRITOS POR USUARIO
func (server *Server) getCarritosByUser(ctx *gin.Context) {

	idParam := ctx.Param("id")

	usuarioID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	usuarioID := int32(usuarioID64)

	carritos, err := server.dbtx.GetCarritosByUserV3(ctx, usuarioID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, c := range carritos {
		response = append(response, gin.H{
			"carrito_id": c.ID,
			"usuario_id": c.UsuarioID,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// GET CARRITOS POR PRODUCTO
func (server *Server) getCarritosByProducto(ctx *gin.Context) {

	idParam := ctx.Param("id")

	productID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	productID := int32(productID64)

	carritos, err := server.dbtx.GetCarritosByProducto(ctx, productID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, c := range carritos {
		response = append(response, gin.H{
			"carrito_id":  c.CarritoID,
			"producto_id": c.ID,
			"nombre":      c.Name,
			"cantidad":    c.Cantidad,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// ADD TO CART
func (server *Server) addToCart(ctx *gin.Context) {

	var req struct {
		CarritoID  int32 `json:"carrito_id"`
		ProductoID int32 `json:"producto_id"`
		Cantidad   int32 `json:"cantidad"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.AddToCart(ctx, dto.AddToCartParams{
		CarritoID:  req.CarritoID,
		ProductoID: req.ProductoID,
		Cantidad:   req.Cantidad,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "producto agregado al carrito 🛒",
	})
}

// GET ITEMS POR USUARIO
func (server *Server) getCartItemsByUser(ctx *gin.Context) {

	idParam := ctx.Param("id")

	var usuarioID int32
	fmt.Sscan(idParam, &usuarioID)

	carritos, err := server.dbtx.GetCarritosByUserV3(ctx, usuarioID)
	if err != nil || len(carritos) == 0 {
		ctx.JSON(http.StatusNotFound, gin.H{"error": "carrito no encontrado"})
		return
	}

	carritoID := carritos[0].ID
	items, err := server.dbtx.GetCartItems(ctx, carritoID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H
	var total float64 = 0

	for _, item := range items {

		precio, _ := strconv.ParseFloat(item.Price, 64)

		subtotal := precio * float64(item.Cantidad)
		total += subtotal

		response = append(response, gin.H{
			"carrito_id": item.CarritoID,
			"producto":   item.Name,
			"precio":     precio,
			"cantidad":   item.Cantidad,
			"subtotal":   subtotal,
		})
	}

	ctx.JSON(http.StatusOK, gin.H{
		"items": response,
		"total": total,
	})
}

// DELETE ITEM
func (server *Server) deleteCartItem(ctx *gin.Context) {

	var req struct {
		CarritoID  int32 `json:"carrito_id"`
		ProductoID int32 `json:"producto_id"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.DeleteCartItem(ctx, dto.DeleteCartItemParams{
		CarritoID:  req.CarritoID,
		ProductoID: req.ProductoID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "producto eliminado del carrito 🗑️",
	})
}
