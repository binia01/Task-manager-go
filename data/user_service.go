package data

import (
	"context"
	"fmt"
	"task-manager-go/middleware"
	"task-manager-go/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var user_collection = db.Database("Task-Database").Collection("Users")

func RegisterUser(user models.User) (bool, error) {
	// Check if username already exists
	var existingUser models.User
	err := user_collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err == nil {
		// Found an existing user with this username
		return false, fmt.Errorf("username already taken")
	} else if err != mongo.ErrNoDocuments {
		// Some other DB error
		return false, fmt.Errorf("error checking existing user: %v", err)
	}

	// Count users to determine role
	count, err := user_collection.CountDocuments(context.TODO(), bson.M{})
	if err != nil {
		return false, fmt.Errorf("error counting users: %v", err)
	}

	if count == 0 {
		user.Role = models.AdminRole
	} else {
		user.Role = models.UserRole
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return false, fmt.Errorf("error hashing password: %v", err)
	}
	user.Password = string(hashedPassword)

	// Insert new user
	user.ID = primitive.NilObjectID // Important: reset ID before insert

	res, err := user_collection.InsertOne(context.TODO(), user)
	if err != nil {
		return false, fmt.Errorf("error inserting user: %v", err)
	}

	if oid, ok := res.InsertedID.(primitive.ObjectID); ok {
		user.ID = oid
	}

	return true, nil
}

func LoginUser(user models.User) (string, error) {
	var existingUser models.User
	err := user_collection.FindOne(context.TODO(), bson.M{"username": user.Username}).Decode(&existingUser)
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	err = bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("invalid username or password")
	}

	claims := jwt.MapClaims{
		"user_id":  existingUser.ID,
		"username": existingUser.Username,
		"role":     existingUser.Role,
		"exp":      time.Now().Add(24 * time.Hour).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	jwtToken, err := token.SignedString(middleware.JwtSecret)
	if err != nil {
		return "", err
	}

	return jwtToken, nil
}

func UpdateRole(username string) (bool, error) {
	filter := bson.M{"username": username}
	update := bson.M{
		"$set": bson.M{
			"role": models.AdminRole,
		},
	}
	_, err := user_collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		return false, fmt.Errorf("can not find user")
	}
	return false, nil
}
