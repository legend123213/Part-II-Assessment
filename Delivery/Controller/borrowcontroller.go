package controller

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	domain "github.com/legend123213/ass2/Domain"
)
type BorrowController struct {
	BorrowUsecase domain.BorrowUsecase
}

	

func NewBorrowController(usecase domain.BorrowUsecase) *BorrowController {
	return &BorrowController{
		BorrowUsecase: usecase,
	}
}

func (controller *BorrowController) CreateBorrow() gin.HandlerFunc{
	return func(c *gin.Context){
		var borrow *domain.Borrow 
		if err := c.ShouldBindJSON(&borrow); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}		
		err := controller.BorrowUsecase.BorrowRequest(borrow.UserEmail,borrow.BookID)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":"the borrow request has been sent",
		})
	}
}
func (controller *BorrowController) GetBorrowRequest() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Query("id")
		email:=c.MustGet("claims").(*domain.Claims).Email
		borrow, err := controller.BorrowUsecase.GetBorrowRequest(id,email)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, borrow)
	}
}
func (controller *BorrowController) GetBorrowRequests() gin.HandlerFunc{
	return func(c *gin.Context){
		borrows, err := controller.BorrowUsecase.GetBorrowRequests()
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, borrows)
	}
}
func (controller *BorrowController) UpdateBorrow() gin.HandlerFunc{
	return func(c *gin.Context){
		var borrow domain.Borrow
		id := c.Param("id")
		if err := c.ShouldBindJSON(&borrow); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(),"sfs":"asdfas"})
			return
		}
		log.Println(borrow.Status)
		if strings.Contains(borrow.Status, "Pending") || strings.Contains(borrow.Status, "Approved") || strings.Contains(borrow.Status, "Rejected") {
			err := controller.BorrowUsecase.UpdateBorrow(id,borrow.Status)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":"the borrow request has been updated",
		})
		}else{
		c.JSON(400, gin.H{
				"error": "Status is required or wronge submitted",
			})
			return
		}
		
	}
}
func (controller *BorrowController) DeleteBorrow() gin.HandlerFunc{
	return func(c *gin.Context){
		id := c.Query("id")
		err := controller.BorrowUsecase.DeleteBorrow(id)
		if err != nil {
			c.JSON(400, gin.H{
				"error": err.Error(),
			})
			return
		}
		c.JSON(200, gin.H{
			"message":"the borrow request has been deleted",
		})
	}
}
