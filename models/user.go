package models

import (
	"encoding/json"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
