package Repositories

import (
	"context"
	"fmt"

	"task-manager-go/Domain"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) Domain.IUserRepository {
	return &userRepository{collection: collection}
}

func (ur *userRepository) FindByUsername(username string) (*Domain.User, error) {
	var user Domain.User
	err := ur.collection.FindOne(context.TODO(), bson.M{"username": username}).Decode(&user)
	if err != nil {
		return nil, fmt.Errorf("user not found")
	}
	return &user, nil
}

func (ur *userRepository) CreateUser(user Domain.User) (*Domain.User, error) {
	// Check if username exists
	var existingUser Domain.User
	err := ur.collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		return nil, fmt.Errorf("username already taken")
	}

	// Count users to determine role
	count, err := ur.collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return nil, fmt.Errorf("error counting users")
	}

	if count == 0 {
		user.Role = Domain.AdminRole
	} else {
		user.Role = Domain.UserRole
	}

	user.ID = primitive.NilObjectID
	res, err := ur.collection.InsertOne(context.TODO(), user)
	if err != nil {
		return nil, fmt.Errorf("error creating user")
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}
	return &user, nil
}

func (ur *userRepository) UpdateUserRole(username string, role Domain.Role) error {
	filter := bson.M{"username": username}
	update := bson.M{"$set": bson.M{"role": role}}

	result, err := ur.collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		return fmt.Errorf("error updating user role")
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
