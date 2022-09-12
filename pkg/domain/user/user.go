package user

import (
	"context"

	"google.golang.org/grpc/metadata"
)

type User struct {
	id        string
	firstname string
	lastname  string
	title     string
}

func NewUser(id string, firstname string, lastname string, title string) *User {
	return &User{
		id:        id,
		firstname: firstname,
		lastname:  lastname,
		title:     title,
	}
}

func (u User) Id() string {
	return u.id
}

func (u User) Firstname() string {
	return u.firstname
}

func (u User) Lastname() string {
	return u.lastname
}

func (u User) Title() string {
	return u.title
}

func FromIncomingContext(ctx context.Context) *User {
	m, ok := metadata.FromIncomingContext(ctx)
	if !ok || m == nil {
		return nil
	}

	return NewUser(
		md(m, "userId"),
		md(m, "userName"),
		md(m, "userLastname"),
		md(m, "userTitle"),
	)
}

func md(m metadata.MD, key string) string {
	c := m.Get(key)
	if c != nil {
		return c[0]
	}
	return ""
}

func FromContext(ctx context.Context) *User {
	c := ctx.Value("user")
	if c != nil {
		u, ok := c.(*User)
		if ok {
			return u
		}
	}
	return nil
}

func NameFromContext(ctx context.Context) string {
	u := FromContext(ctx)
	return u.firstname
}

func ContextWithValue(ctx context.Context) context.Context {
	return context.WithValue(ctx, "user", FromIncomingContext(ctx))
}

func NewContextWithValue(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), "user", FromContext(ctx))
}

func NewContextWithValueFromIncomingContext(ctx context.Context) context.Context {
	return context.WithValue(context.Background(), "user", FromIncomingContext(ctx))
}
