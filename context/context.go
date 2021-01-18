package context

import (
	"context"
)

const (
	userKey privateKey = "user"
)

type privateKey string

func WithUser(ctx context.Context, user *models.User) context.Context {
	return context.WithValue(ctx, privateKey, user)
}

func User(ctx context.Context) *User {
	if temp := ctx.Value(privateKey); temp != nil {
		if user, ok := temp.(*models.User); ok {
			return user
		}
	}
	return nil
}
