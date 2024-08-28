package routes

// gin-gonic
import (
	"github.com/gin-gonic/gin"
	controller "github.com/legend123213/ass2/Delivery/Controller"
	infrastructure "github.com/legend123213/ass2/Infrastructure"
)
type Route struct {
	UserController *controller.UserController 
	BorrowController controller.BorrowController
}

func NewRoute(userController *controller.UserController,borrowController controller.BorrowController) *Route {
	return &Route{
		UserController: userController,
		BorrowController: borrowController,
	}
}

func (route *Route) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/users", route.UserController.CreateUser())
	r.GET("/users/", route.UserController.VerifyUser())
	r.POST("/users/login", route.UserController.Login())
	r.GET("/users/password-reset", route.UserController.ResetPasswordRequest())
	r.PUT("/users/password-update", route.UserController.UpdatePassword())
	r.Use(infrastructure.AuthenticationMiddleware())
	r.POST("/books/borrow",route.BorrowController.CreateBorrow())
	r.GET("/books/borrow",route.BorrowController.GetBorrowRequest())
	r.Use(infrastructure.ADMINMiddleware())
	r.PATCH("/books/borrow/:id/status",route.BorrowController.UpdateBorrow())
	r.DELETE("/books/borrow/:id",route.BorrowController.DeleteBorrow())
	r.GET("/admin/borrows",route.BorrowController.GetBorrowRequests())
	r.DELETE("/users/:id", route.UserController.DeleteUser())
	return r
}



