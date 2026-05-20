package api

import (
	"database/sql"
	"net/http"
	"strconv"
	"time"

	"rest/dto"

	"github.com/gin-gonic/gin"
)

type loginRequest struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type LoginResponse struct {
	AccessToken string  `json:"access_token"`
	Payload     payload `json:"payload"`
}

type payload struct {
	ID   int32  `json:"id"`
	Name string `json:"name"`
	Role string `json:"role"`
}

func (server *Server) login(ctx *gin.Context) {
	var req loginRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	user, err := server.dbtx.GetUserByEmail(ctx, sql.NullString{
		String: req.Email,
		Valid:  true,
	})

	if err != nil {
		if err == sql.ErrNoRows {
			ctx.JSON(http.StatusNotFound, errorResponse(err))
			return
		}
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	if user.Password.String != req.Password {
		ctx.JSON(http.StatusUnauthorized, gin.H{"message": "Usuario no autorizado"})
		return
	}

	accessToken, err := server.tokenBuilder.CreateToken(
		user.Role.String,
		user.Name,
		"",
		time.Minute*60,
	)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	resp := LoginResponse{
		AccessToken: accessToken,
		Payload: payload{
			ID:   user.ID,
			Name: user.Name,
			Role: user.Role.String,
		},
	}

	ctx.JSON(http.StatusOK, resp)
}

// GET ALL USuarios
func (server *Server) getAllUsers(ctx *gin.Context) {

	users, err := server.dbtx.GetAllUsers(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	var response []gin.H

	for _, u := range users {

		response = append(response, gin.H{
			"id":    u.ID,
			"name":  u.Name,
			"email": u.Email.String,
			"role":  u.Role.String,
		})
	}

	ctx.JSON(http.StatusOK, response)
}

// GET USUARIO BY ID
func (server *Server) getUserById(ctx *gin.Context) {

	idParam := ctx.Param("id")

	userID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := int32(userID64)

	user, err := server.dbtx.GetUserById(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email.String,
		"role":  user.Role.String,
	})
}

// UPDATE USUARIO
func (server *Server) updateUser(ctx *gin.Context) {

	var req struct {
		ID    int32  `json:"id"`
		Name  string `json:"name"`
		Email string `json:"email"`
	}

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	err := server.dbtx.UpdateUser(ctx, dto.UpdateUserParams{
		Name: req.Name,

		Email: sql.NullString{
			String: req.Email,
			Valid:  true,
		},

		ID: req.ID,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuario actualizado 👤",
	})
}

// DELETE USER
func (server *Server) deleteUser(ctx *gin.Context) {

	idParam := ctx.Param("id")

	userID64, err := strconv.ParseInt(idParam, 10, 32)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, errorResponse(err))
		return
	}

	userID := int32(userID64)

	err = server.dbtx.DeleteUser(ctx, userID)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"message": "usuario eliminado 🗑️",
	})
}

// GET USER BY EMAIL
func (server *Server) getUserByEmail(ctx *gin.Context) {

	email := ctx.Param("email")

	user, err := server.dbtx.GetUserByEmail(ctx, sql.NullString{
		String: email,
		Valid:  true,
	})

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, errorResponse(err))
		return
	}

	ctx.JSON(http.StatusOK, gin.H{
		"id":    user.ID,
		"name":  user.Name,
		"email": user.Email.String,
		"role":  user.Role.String,
	})
}
