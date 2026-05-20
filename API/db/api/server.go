package api

import (
	"rest/db/security"
	"rest/dto"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type Server struct {
	dbtx         *dto.Queries
	router       *gin.Engine
	tokenBuilder security.Builder
}

func NewServer(dbtx *dto.Queries, secret string) (*Server, error) {

	builder, err := security.NewPasetoBuilder(secret)
	if err != nil {
		return nil, err
	}

	server := &Server{
		dbtx:         dbtx,
		tokenBuilder: builder,
	}

	router := gin.Default()

	router.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: false,
		MaxAge:           50 * time.Second,
	}))

	router.POST("/api/v1/login", server.login)
	router.GET("/api/v1/category", server.getAll)
	router.GET("/api/v1/category/:id", server.getCategoryById)
	router.GET("/api/v1/products", server.getAllProducts)
	router.GET("/api/v1/products/category/:id", server.getProductsByCategory)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenBuilder))

	authRoutes.GET("/api/v1/users", server.getAllUsers)
	authRoutes.GET("/api/v1/users/:id", server.getUserById)
	authRoutes.GET("/api/v1/users/email/:email", server.getUserByEmail)
	authRoutes.PUT("/api/v1/users", server.updateUser)
	authRoutes.DELETE("/api/v1/users/:id", server.deleteUser)
	authRoutes.POST("/api/v1/category", server.createCategory)
	authRoutes.PUT("/api/v1/category", server.updateCategory)
	authRoutes.DELETE("/api/v1/category/:id", server.deleteCategory)
	authRoutes.POST("/api/v1/products", server.createProduct)
	authRoutes.PUT("/api/v1/products/:id", server.updateProduct)
	authRoutes.DELETE("/api/v1/products", server.deleteProduct)
	authRoutes.GET("/api/v1/carritos", server.getAllCarritos)
	authRoutes.GET("/api/v1/carritos/producto/:id", server.getCarritosByProducto)
	authRoutes.GET("/api/v1/carritos/user/:id", server.getCarritosByUser)
	authRoutes.GET("/api/v1/carritos/user/:id/items", server.getCartItemsByUser)
	authRoutes.POST("/api/v1/cart/create", server.createCarrito)
	authRoutes.POST("/api/v1/cart/add", server.addToCart)
	authRoutes.DELETE("/api/v1/cart/item", server.deleteCartItem)
	authRoutes.GET("/api/v1/compra", server.getAllCompras)
	authRoutes.GET("/api/v1/compra/:id", server.getDetalleCompra)
	authRoutes.GET("/api/v1/compra/user/:id", server.getComprasByUser)
	authRoutes.POST("/api/v1/compra/create", server.createCompra)
	authRoutes.POST("/api/v1/compra/detail", server.createDetalleCompra)

	server.router = router

	return server, nil
}

func (server *Server) Start(url string) error {
	return server.router.Run(url)
}

func errorResponse(err error) gin.H {
	return gin.H{
		"error": err.Error(),
	}
}
