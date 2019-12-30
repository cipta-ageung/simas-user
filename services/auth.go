package services

import (
	"context"

	auth "github.com/cipta-ageung/simas-user/proto/auth"
)

// AuthService : struct interface
type AuthService struct{}

// Auth : method get authentication
func (g *AuthService) Auth(ctx context.Context, req *auth.User, rsp *auth.Token) error {
	return nil
}

// Create : method create user
func (g *AuthService) Create(ctx context.Context, req *auth.User, rsp *auth.Response) error {
	return nil
}

// Delete : method delete user
func (g *AuthService) Delete(ctx context.Context, req *auth.User, rsp *auth.Response) error {
	return nil
}

// Get : method get user
func (g *AuthService) Get(ctx context.Context, req *auth.User, rsp *auth.Response) error {
	return nil
}

// GetAll : method get all user
func (g *AuthService) GetAll(ctx context.Context, req *auth.Request, rsp *auth.Response) error {
	return nil
}

// ValidateToken : method validation token if expired token
func (g *AuthService) ValidateToken(ctx context.Context, req *auth.Token, rsp *auth.Token) error {
	return nil
}
