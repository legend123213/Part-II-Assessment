package routes

// gin-gonic
import (
	"github.com/gin-gonic/gin"
	controller "github.com/legend123213/ass2/Delivery/Controller"
)
type Route struct {
	UserController *controller.UserController
}

func NewRoute(userController *controller.UserController) *Route {
	return &Route{
		UserController: userController,
	}
}

func (route *Route) SetupRouter() *gin.Engine {
	r := gin.Default()
	r.Use(gin.Recovery())
	r.POST("/users", route.UserController.CreateUser())
	r.GET("/users/", route.UserController.VerifyUser())
	r.GET("/users/password-reset", route.UserController.ResetPasswordRequest())
	r.PUT("/users/password-update", route.UserController.UpdatePassword())
	r.DELETE("/users/:id", route.UserController.DeleteUser())
	r.POST("/users/login", route.UserController.Login())
	return r
}



