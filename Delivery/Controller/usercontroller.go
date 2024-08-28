package controller

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	domain "github.com/legend123213/ass2/Domain"
	usecase "github.com/legend123213/ass2/Usecase"
)

type UserController struct {
	UserUsecase usecase.UserServiceUsecase
}

	

func NewUserController(usecase usecase.UserServiceUsecase) *UserController {
	return &UserController{
		UserUsecase: usecase,
	}
}

func (controller *UserController) CreateUser() gin.HandlerFunc{
	return func(c *gin.Context){
		var user *domain.User 
		if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}		
		err := controller.UserUsecase.SendVerifyRegisterationUser(user)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":"the activation link has been sent to your email",
		})
	}
}
func (controller *UserController) VerifyUser() gin.HandlerFunc{
	return func(c *gin.Context){
		email := c.Query("email")
		token := c.Query("token")
		log.Println(email,token,"asdf")
		_, err := controller.UserUsecase.VerifyRegisterationUser(email,token)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "your account has been verified you can login now",
		})
	}
}


func (controller *UserController) GetUser() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		user, err := controller.UserUsecase.GetUser(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"user": user,
		})
	}
}

func (controller *UserController) ResetPasswordRequest() gin.HandlerFunc{
	return func(c *gin.Context){
		email := c.MustGet("claims").(*domain.Claims).Email
		err := controller.UserUsecase.UpdatePasswordRequest(email)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Reset password link sent to your email",
		})
	}
}
func (controller *UserController) UpdatePassword() gin.HandlerFunc{
	return func(c *gin.Context){
	  var forgetPassword *domain.ForgetPassword
	  forgetPassword.Email = c.Query("email")
		forgetPassword.VerifyToken = c.Query("token")
		if err := c.ShouldBindJSON(&forgetPassword); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if forgetPassword.Password != forgetPassword.ConfrimPassword {
			c.JSON(400, gin.H{
				"error": "Password and Confirm Password does not match",
			})
			return
		}
		err := controller.UserUsecase.ResetPassword(*forgetPassword)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "Password updated successfully",
		})

}}

func (controller *UserController) DeleteUser() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Param("id")
		err := controller.UserUsecase.DeleteUser(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message": "User deleted successfully",
		})
	}
}

func (controller *UserController) FatchUsers() gin.HandlerFunc{
	return func(c *gin.Context){
		users, err := controller.UserUsecase.GetUsers()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"users": users,
		})
	}
}
func (controller *UserController) Login() gin.HandlerFunc{
	return func(c *gin.Context){
		var user *domain.User
		if err := c.ShouldBindJSON(&user); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		accessToken,refreshToken, err := controller.UserUsecase.LoginUser(user.Email, user.Password)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"accesstoken": accessToken,
			"refreshtoken": refreshToken,
		})
	}
}
