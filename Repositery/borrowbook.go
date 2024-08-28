package repositery

import (
	"context"

	domain "github.com/legend123213/ass2/Domain"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)





type BookServiceRepo struct{
	DBClient *mongo.Database

}

func NewBookServiceRepo(client *mongo.Database) *UserServiceRepo{
	return &UserServiceRepo{
		DBClient: client,
	}
}

func (repo *UserServiceRepo) CreateBorrow(borrow domain.Borrow) error{
	borrow.ID = primitive.NewObjectID().Hex()
	collection := repo.DBClient.Collection("Requests")
	_, err := collection.InsertOne(context.Background(), borrow)
	if err != nil {
		return err
	}
	return nil
}
func (repo *UserServiceRepo) GetBorrow(id string,email string) (*domain.Borrow, error){
	collection := repo.DBClient.Collection("Requests")
	var borrow domain.Borrow
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}
	err = collection.FindOne(context.Background(), bson.M{"_id": oid, "user_email": email}).Decode(&borrow)
	if err != nil {
		return nil, err
	}
	return &borrow, nil
}
func (repo *UserServiceRepo) GetBorrowRequests() ([]*domain.Borrow, error){
	collection := repo.DBClient.Collection("Requests")
	cursor, err := collection.Find(context.Background(), bson.D{})
	if err != nil {
		return nil, err
	}
	var borrows []*domain.Borrow
	for cursor.Next(context.Background()) {
		var borrow domain.Borrow
		cursor.Decode(&borrow)
		borrows = append(borrows, &borrow)
	}
	return borrows, nil
}
func (repo *UserServiceRepo) UpdateBorrow(id string, status string) error{
	collection := repo.DBClient.Collection("Requests")
	var borrow domain.Borrow
	update := bson.M{
		"$set": bson.M{
			"status":       status,
		},
	}
	if err := collection.FindOneAndUpdate(context.TODO(), bson.M{"_id": id}, update,options.FindOneAndUpdate().SetReturnDocument(1)).Decode(&borrow); err!=nil{
		return err
	}
	return nil
}
func (repo *UserServiceRepo) DeleteBorrow(id string) error{
	collection := repo.DBClient.Collection("Requests")
	oid, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.DeleteOne(context.Background(), bson.M{"_id": oid})
	if err != nil {
		return err
	}
	return nil
}
