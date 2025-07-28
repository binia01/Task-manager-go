package Domain

import (
	"encoding/json"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          string    `json:"id"`
	Title       string    `json:"title"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Status      string    `json:"status"`
}

type User struct {
	ID       primitive.ObjectID `json:"_id" bson:"_id,omitempty"`
	Username string             `json:"username"`
	Password string             `json:"-"` // hidden in JSON
	Role     Role               `json:"role"`
}

type Role int

const (
	UserRole Role = iota
	AdminRole
)

func (r Role) String() string {
	switch r {
	case UserRole:
		return "user"
	case AdminRole:
		return "Admin"
	default:
		return "unknown"
	}
}

func (r Role) MarshalJSON() ([]byte, error) {
	return json.Marshal(r.String())
}

type ITaskRepository interface {
	GetAllTasks() ([]Task, error)
	GetTaskById(id string) (*Task, error)
	CreateTask(task Task) error
	UpdateTask(id string, task Task) (*Task, error)
	DeleteTask(id string) (*Task, error)
}

type IUserRepository interface {
	FindByUsername(username string) (*User, error)
	CreateUser(user User) (*User, error)
	UpdateUserRole(username string, role Role) error
}

type IPasswordService interface {
	HashPassword(password string) (string, error)
	ComparePassword(hashedPassword string, password string) error
}

type IJWTService interface {
	GenerateToken(user User) (string, error)
}
