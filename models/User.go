package models

import (
	"context"

	"github.com/google/uuid"
)

type ContextKey string

const UserContext ContextKey = "userContext"

type User struct {
	ID   uuid.UUID
	Name string
}

func GetUser(ctx context.Context) *User {
	if u, ok := ctx.Value(UserContext).(*User); ok {
		return u
	}

	return NewUser("Default User")
}

func NewUser(n string) *User {
	return &User{
		ID:   uuid.New(),
		Name: n,
	}
}
