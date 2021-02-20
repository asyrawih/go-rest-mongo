package repository

import (
	"github.com/hananloser/rest-fiber-mongo-2/config"
	"github.com/hananloser/rest-fiber-mongo-2/entity"
	"github.com/hananloser/rest-fiber-mongo-2/exception"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepositoryImpl struct {
	Collection *mongo.Collection
}

func (u UserRepositoryImpl) Insert(user entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()

	_, err := u.Collection.InsertOne(ctx, bson.M{
		"_id":          user.Id,
		"name":         user.Name,
		"email":        user.Email,
		"password":     user.Password,
		"address":      user.Address,
		"phone_number": user.PhoneNumber,
	})

	exception.PanicIfNeeded(err)
}

// Get All User collection
func (u UserRepositoryImpl) GetAll(userId string) (users []entity.User) {
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	query := bson.M{}

	if userId != "" {
		query = bson.M{"_id" : userId}
	}

	cursor, err := u.Collection.Find(ctx, query)
	exception.PanicIfNeeded(err)

	var documents []bson.M
	err = cursor.All(ctx, &documents)

	for _, document := range documents {
		users = append(users, entity.User{
			Id:          document["_id"].(string),
			Name:        document["name"].(string),
			Email:       document["email"].(string),
			Password:    document["password"].(string),
			Address:     document["address"].(string),
			PhoneNumber: document["phone_number"].(string),
		})
	}

	return users
}

func (u UserRepositoryImpl) Find(Id string) (users entity.User) {
	filter := bson.D{{"_id", Id}}
	ctx, cancel := config.NewMongoContext()
	defer cancel()
	err := u.Collection.FindOne(ctx, filter).Decode(&users)
	exception.PanicIfNeeded(err)
	return users
}

func (u UserRepositoryImpl) DeleteAll() {
	panic("implement me")
}

func NewUserRepository(database *mongo.Database) UserRepository {
	return &UserRepositoryImpl{
		Collection: database.Collection("users"),
	}
}
