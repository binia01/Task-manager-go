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
