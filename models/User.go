package models

import "context"

type ContextKey string

const UserContext ContextKey = "userContext"

type User struct {
	ID   int
	Name string
}

func GetUser(ctx context.Context) User {
	if u, ok := ctx.Value(UserContext).(User); ok {
		return u
	}

	return User{
		ID:   -1,
		Name: "User",
	}
}
