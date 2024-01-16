// package api

// import (
// 	"fmt"
// 	db "new/learning/user/db/sqlc"
// 	"new/learning/user/token"

// 	"github.com/gin-gonic/gin"

// )

// // server serves http reqests for our banking service
// type Server struct {
// 	store  *db.Store
// 	tokenMaker token.Maker
// 	router *gin.Engine
// }

// // NewServer creates a new HTTP server and set up routing.
// func NewServer(config util.Config, store *db.Store) (*Server,error){

// 	tokenMaker,err := token.NewPasetoMaker("")

// 	if err != nil{
// 		return nil, fmt.Errorf("cannot create token maker")
// 	}

// 	server := &Server{
// 		store: store,
// 		tokenMaker: tokenMaker,
// 	}

// 	router := gin.Default()
// 	router.POST("/users", server.createUser)
// 	router.POST("/accounts", server.createAccount)
// 	router.GET("accounts/:id", server.getAccount)
// 	// router.GET("accounts/", server.ListAccount)
// 	router.POST("/transfers", server.createTransfer)

// 	server.router = router
// 	return server, err
// }

// // Start runs the HTTP server on a specific address.
// func (server *Server) Start(address string) error {
// 	return server.router.Run(address)
// }

// func errorResponse(err error) gin.H {
// 	return gin.H{"error": err.Error()}
// }

package api

import (
	"fmt"
	db "new/learning/user/db/sqlc"
	"new/learning/user/token"
	"new/learning/user/util" // Import the util package

	"github.com/gin-gonic/gin"
)

// Server serves HTTP requests for our banking service
type Server struct {
	config     util.Config
	store      *db.Store
	tokenMaker token.Maker
	router     *gin.Engine
}

// NewServer creates a new HTTP server and sets up routing.
func NewServer(config util.Config, store *db.Store) (*Server, error) {
	tokenMaker, err := token.NewPasetoMaker(config.TokenSymmetricKey) // Use the token symmetric key from config

	if err != nil {
		return nil, fmt.Errorf("cannot create token maker: %v", err)
	}

	server := &Server{
		config:     config,
		store:      store,
		tokenMaker: tokenMaker,
	}

	// router := gin.Default()
	// router.POST("/users", server.createUser)
	// router.POST("/users/login", server.loginUser)

	// router.POST("/accounts", server.createAccount)
	// router.GET("/accounts/:id", server.getAccount) // Fixed missing leading slash
	// // router.GET("/accounts/", server.ListAccount) // Uncomment this line if needed
	// router.POST("/transfers", server.createTransfer)

	// server.router = router
	server.setupRouter()
	return server, nil // Return server with nil error as we've handled error in tokenMaker creation
}

func (server *Server) setupRouter(){

	router := gin.Default()
	router.POST("/users", server.createUser)
	router.POST("/users/login", server.loginUser)

	authRoutes := router.Group("/").Use(authMiddleware(server.tokenMaker))

	authRoutes.POST("/accounts", server.createAccount)
	authRoutes.GET("/accounts/:id", server.getAccount) // Fixed missing leading slash
	// router.GET("/accounts/", server.ListAccount) // Uncomment this line if needed
	authRoutes.POST("/transfers", server.createTransfer)

	server.router = router
}

// Start runs the HTTP server on a specific address.
func (server *Server) Start(address string) error {
	return server.router.Run(address)
}

func errorResponse(err error) gin.H {
	return gin.H{"error": err.Error()}
}
