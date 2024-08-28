package usecase

import (
	domain "github.com/legend123213/ass2/Domain"
)



type BorrowServiceUsecase struct{
	BorrowRepo domain.BorrowService
}

func NewBorrowServiceUsecase(repo domain.BorrowService) *BorrowServiceUsecase{
	return &BorrowServiceUsecase{
		BorrowRepo: repo,
	}
}
func (usecase *BorrowServiceUsecase) BorrowRequest(email string, bookid string) error{
var borrow domain.Borrow
borrow.BookID = bookid
borrow.UserEmail = email
borrow.Status = "Pending"
	return usecase.BorrowRepo.CreateBorrow(borrow)
}
func (usecase *BorrowServiceUsecase) GetBorrowRequest(id string,email string) (*domain.Borrow, error){
	return usecase.BorrowRepo.GetBorrow(id,email)
}
func (usecase *BorrowServiceUsecase) GetBorrowRequests() ([]*domain.Borrow, error){
	return usecase.BorrowRepo.GetBorrowRequests()
}
func (usecase *BorrowServiceUsecase) UpdateBorrow(id string, status string) error{
	return usecase.BorrowRepo.UpdateBorrow(id,status)
}
func (usecase *BorrowServiceUsecase) DeleteBorrow(id string) error{
	return usecase.BorrowRepo.DeleteBorrow(id)
}