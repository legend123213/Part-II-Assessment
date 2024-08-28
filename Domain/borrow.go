package domain

import "github.com/gin-gonic/gin"
type Borrow struct{
	ID string `json:"id" bson:"_id"`
	BookID string `json:"book_id" bson:"book_id"`
	UserEmail string `json:"user_email" bson:"user_email"`
	Status string `json:"status" bson:"status"`
 }

type BorrowService interface{
	CreateBorrow(borrow Borrow) error
	GetBorrow(id string,email string) (*Borrow, error)
	GetBorrowRequests() ([]*Borrow, error)
	UpdateBorrow(id string, status string) error
	DeleteBorrow(id string) error
}
type BorrowUsecase interface{
	BorrowRequest(email string, bookid string) error
	GetBorrowRequest(id string,email string) (*Borrow, error)
	GetBorrowRequests() ([]*Borrow, error)
	UpdateBorrow(id string, status string) error
	DeleteBorrow(id string) error
}
type BorrowControllerinter interface
{
	CreateBorrow() gin.HandlerFunc
	GetBorrowRequest() gin.HandlerFunc
	GetBorrowRequests() gin.HandlerFunc
	UpdateBorrow() gin.HandlerFunc
	DeleteBorrow() gin.HandlerFunc
}