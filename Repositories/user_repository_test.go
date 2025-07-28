package Repositories_test

import (
	"context"
	"testing"

	"task-manager-go/Domain"
	"task-manager-go/Repositories"

	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func setupUserTestDB() *mongo.Collection {
	client, _ := mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb+srv://username:password@cluster0.tj8um.mongodb.net/?retryWrites=true&w=majority"))
	col := client.Database("task_test_db").Collection("users")
	col.DeleteMany(context.TODO(), map[string]interface{}{})
	return col
}

func TestUserRepository_CreateAndFind(t *testing.T) {
	col := setupUserTestDB()
	repo := Repositories.NewUserRepository(col)

	user := Domain.User{ID: primitive.NilObjectID, Username: "testuser", Password: "pass"}
	created, err := repo.CreateUser(user)
	assert.NoError(t, err)
	assert.Equal(t, "testuser", created.Username)

	found, err := repo.FindByUsername("testuser")
	assert.NoError(t, err)
	assert.Equal(t, created.Username, found.Username)
}

func TestUserRepository_UpdateRole(t *testing.T) {
	col := setupUserTestDB()
	repo := Repositories.NewUserRepository(col)

	user := Domain.User{ID: primitive.NilObjectID, Username: "roleuser", Password: "pw"}
	_, _ = repo.CreateUser(user)

	err := repo.UpdateUserRole("roleuser", Domain.AdminRole)
	assert.NoError(t, err)

	updated, _ := repo.FindByUsername("roleuser")
	assert.Equal(t, Domain.AdminRole, updated.Role)
}
