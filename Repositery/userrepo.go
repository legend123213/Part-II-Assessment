package repositery

import (
	"context"
	"errors"

	domain "github.com/legend123213/ass2/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)





type UserServiceRepo struct{
	DBClient *mongo.Database

}

func NewUserServiceRepo(client *mongo.Database) *UserServiceRepo{
	return &UserServiceRepo{
		DBClient: client,
	}
}

func (repo *UserServiceRepo) CreateUser(user *domain.User) error{
	user.ID = primitive.NewObjectID().Hex()
	collection := repo.DBClient.Collection("users")
	_, err := collection.InsertOne(context.Background(), user)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserServiceRepo) GetUser(id string) (*domain.User, error){
	collection := repo.DBClient.Collection("users")
	var user domain.User
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.Background(), bson.E{"_id",oid}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}
func (repo *UserServiceRepo) GetUserByEmail(email string) (domain.User, error){
	var user domain.User
	err := repo.DBClient.Collection("users").FindOne(context.TODO(), bson.D{{"email", email}}).Decode(&user)
	if err != nil {
		return domain.User{}, err
	}
	return user, nil
}
func (repo *UserServiceRepo) GetUsers() ([]*domain.User, error){
	collection := repo.DBClient.Collection("users")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var users []*domain.User
	for cursor.Next(context.Background()){
		var user domain.User
		cursor.Decode(&user)
		users = append(users, &user)
	}
	return users, nil
}

func (repo *UserServiceRepo) MakeAcitiveUser(user domain.User) (domain.User, error){
	update := bson.M{
		"$set": bson.M{
			"is_active":       user.IsActive,
			"verify_token":    user.VerifyToken,
		},
	}
	collection := repo.DBClient.Collection("users")
	if err := collection.FindOneAndUpdate(context.TODO(), bson.M{"email": user.Email}, update,options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user); err!=nil{
		return user,errors.New("")
	}
	
	return user,nil
}
func (repo *UserServiceRepo) UpdateUsertoken(user *domain.User) error{
	update := bson.M{
		"$set": bson.M{
			"verify_token":    user.VerifyToken,
			"expiration_date_token": user.ExpirationDatetoken,
		},
	}
	collection := repo.DBClient.Collection("users")
	if err := collection.FindOneAndUpdate(context.TODO(), bson.M{"email": user.Email}, update,options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user); err!=nil{
		return errors.New("")
	}
	return nil
}
func (repo *UserServiceRepo)UpdatePassword(user *domain.User) error{
	update := bson.M{
		"$set": bson.M{
			"password": user.Password,
		},
	}
	collection := repo.DBClient.Collection("users")
	if err := collection.FindOneAndUpdate(context.TODO(), bson.M{"email": user.Email}, update,options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&user); err!=nil{
		return errors.New("")
	}
	return nil
}
func (repo *UserServiceRepo) DeleteUser(id string) error{
	collection := repo.DBClient.Collection("users")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.E{"_id", oid})
	if err != nil {
		return err
	}
	return nil
}
