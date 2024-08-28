package usecase

import (
	"crypto/rand"
	"errors"
	"log"
	"math/big"
	"time"

	domain "github.com/legend123213/ass2/Domain"
	infrastructure "github.com/legend123213/ass2/Infrastructure"
)



type UserServiceUsecase struct{
	UserRepo domain.UserRepo
}

func NewUserServiceUsecase(repo domain.UserRepo) *UserServiceUsecase{
	return &UserServiceUsecase{
		UserRepo: repo,
	}
}	
func (usecase *UserServiceUsecase) SendVerifyRegisterationUser(user *domain.User)  error {
	user.ID = ""
	user.Password,_=infrastructure.PasswordHasher(user.Password)
	users,err := usecase.UserRepo.GetUsers()

	if err != nil {
		return  err
	}
	if len(users) ==0 {
		user.IsAdmin = true
	}else{
		user.IsAdmin = false
	}
	user.IsActive = false
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	confirmationToken := make([]byte, 64)
	charsetLength := big.NewInt(int64(len(charset)))
	for i := 0; i < 64; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		confirmationToken[i] = charset[num.Int64()]
	}
	user.VerifyToken = string(confirmationToken)

	user.ExpirationDatetoken = time.Now().Add(2 * time.Hour)
	
	link := "http://localhost:8000/users/activationusers/?email=" + user.Email + "&token=" + string(confirmationToken)
	err = infrastructure.SendEmail(user.Email, "Verify your Account", "Please verify your email by clicking the following link: ", link )
	if err != nil {
		log.Println(err,"sdfasd")
		return err
	}
	
	return usecase.UserRepo.CreateUser(user)
}
func (usecase *UserServiceUsecase) VerifyRegisterationUser(email string,VerifyToken string) (domain.User, error){
	user, err := usecase.UserRepo.GetUserByEmail(email)
	if err != nil {
		return domain.User{}, err
	}
	if user.VerifyToken != VerifyToken {
		return domain.User{}, errors.New("Invalid Token")
	}
	if user.ExpirationDatetoken.Before(time.Now()) {
		return domain.User{}, errors.New("Token Expired")
	}

	user.IsActive = true
	user.VerifyToken = ""
	user.ExpirationDatetoken = time.Time{}
	return usecase.UserRepo.MakeAcitiveUser(user)

}
func (usecase *UserServiceUsecase) GetUser(id string) (*domain.User, error){
	return usecase.UserRepo.GetUser(id)
}
func (usecase *UserServiceUsecase) GetUsers() ([]*domain.User, error){
	return usecase.UserRepo.GetUsers()
}


func (usecase *UserServiceUsecase) UpdatePasswordRequest(email string)  error{
	user, err := usecase.UserRepo.GetUserByEmail(email)
	if err != nil {
		return err
	}
	const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	token := make([]byte, 8)
	charsetLength := big.NewInt(int64(len(charset)))
	for i := 0; i < 8; i++ {
		num, _ := rand.Int(rand.Reader, charsetLength)
		token[i] = charset[num.Int64()]
	}
	user.VerifyToken=(string(token))
	user.ExpirationDatetoken = time.Now().Add(2 * time.Hour)
	link := "http://localhost:8000/users/resetpassword/?email=" + user.Email + "&token=" + string(token)
	err = infrastructure.SendEmail(user.Email, "Reset Password", "Please reset your password by clicking the following link: ","<a href=\"" + link + "\">Reset Password</a>")

	if err != nil {
		return  err
	}
	err = usecase.UserRepo.UpdateUsertoken(&user)
	if err != nil {
		return err
	}
	return nil
}
func (usecase *UserServiceUsecase) ResetPassword(userdata domain.ForgetPassword) error{
	user, err := usecase.UserRepo.GetUserByEmail(userdata.Email)
	if err != nil {
		return err
	}
	if user.VerifyToken != userdata.VerifyToken {
		return errors.New("Invalid Token")
	}
	if user.ExpirationDatetoken.Before(time.Now()) {
		return errors.New("Token Expired")
	}

	user.Password,_=infrastructure.PasswordHasher(userdata.Password)
	user.VerifyToken = ""
	user.ExpirationDatetoken = time.Time{}
	return usecase.UserRepo.UpdatePassword(&user)
}
func (usecase *UserServiceUsecase) DeleteUser(id string) error{
	return usecase.UserRepo.DeleteUser(id)
}
func (usecase *UserServiceUsecase) LoginUser(email string, password string) (string,string, error){
	user, err := usecase.UserRepo.GetUserByEmail(email)
	if err != nil {
		return "", "", err
	}
	if user.IsActive == false {
		return "", "", errors.New("User is not active")
	}
	accessToken, refreshToken, err := infrastructure.GenerateToken(&user, password)
	if err != nil {
		return "", "", err
	}

	return accessToken,refreshToken, nil
}