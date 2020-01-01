package services

import (
	"context"

	auth "github.com/cipta-ageung/simas-user/proto/auth"
	empty "github.com/golang/protobuf/ptypes/empty"
)

// AuthService : struct interface
type AuthService struct{}

// Auth : method get authentication
func (g *AuthService) Auth(ctx context.Context, req *auth.User, rsp *auth.Token) error {
	return nil
}

// Create : method create user
func (g *AuthService) Create(ctx context.Context, req *auth.CreateUserRequest, rsp *auth.CreateUserResponse) error {
	return nil
}

// Delete : method delete user
func (g *AuthService) Delete(ctx context.Context, req *auth.DeleteUserRequest, rsp *auth.DeleteUserResponse) error {
	return nil
}

// Read : method get user
func (g *AuthService) Read(context.Context, *auth.ReadUserRequest, *auth.ReadUserResponse) error {
	return nil
}

// Update : method update user
func (g *AuthService) Update(context.Context, *auth.UpdateUserRequest, *auth.UpdateUserResponse) error {
	return nil
}

// UpdateSet : method update list user
func (g *AuthService) UpdateSet(context.Context, *auth.UpdateSetUserRequest, *auth.UpdateSetUserResponse) error {
	return nil
}

// List : method get all user
func (g *AuthService) List(context.Context, *auth.ListUserRequest, *auth.ListUserResponse) error {
	return nil
}

// CustomMethod : custom method
func (g *AuthService) CustomMethod(context.Context, *empty.Empty, *empty.Empty) error {
	return nil
}

// ValidateToken : method validation token if expired token
func (g *AuthService) ValidateToken(ctx context.Context, req *auth.Token, rsp *auth.Token) error {
	return nil
}
