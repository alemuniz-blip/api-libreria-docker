package api

import (
	"net/http"
	"strconv"

	"rest/dto"

	"github.com/gin-gonic/gin"
)

func (server *Server) createCategory(ctx *gin.Context) {

	var req struct {
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	result, err := server.dbtx.CreateCategory(ctx, req.Name)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	lastId, _ := result.LastInsertId()

	ctx.JSON(http.StatusOK, gin.H{
		"generated_id": lastId,
	})
}
func (server *Server) getAll(ctx *gin.Context) {
	categories, err := server.dbtx.GetAllCategories(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}
	ctx.JSON(http.StatusOK, categories)

}

// GET CATEGORY por ID
func (server *Server) getCategoryById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	category, err := server.dbtx.GetCategoryById(ctx, int32(toInt(idParam)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, category)
}

// UPDATE CATEGORY
func (server *Server) updateCategory(ctx *gin.Context) {

	var req struct {
		ID   int32  `json:"id"`
		Name string `json:"name"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.UpdateCategory(ctx, dto.UpdateCategoryParams{
		ID:   req.ID,
		Name: req.Name,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "categoria actualizada 📚",
	})
}

// DELETE CATEGORY
func (server *Server) deleteCategory(ctx *gin.Context) {

	idParam := ctx.Param("id")

	err := server.dbtx.DeleteCategory(ctx, int32(toInt(idParam)))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "categoria eliminada 🗑️",
	})
}

func toInt(value string) int {
	num, _ := strconv.Atoi(value)
	return num
}
