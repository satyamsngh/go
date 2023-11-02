package services

import (
	"context"
	"fmt"
	"job-portal-api/internal/models"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// Conn is our main struct, including the database instance for working with data.

// CreateUser is a method that creates a new user record in the database.
func (s *Store) CreateUser(ctx context.Context, nu models.NewUser) (models.User, error) {

	// We hash the user's password for storage in the database.
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(nu.Password), bcrypt.DefaultCost)
	if err != nil {
		return models.User{}, fmt.Errorf("generating password hash: %w", err)
	}

	// We prepare the User record.
	u := models.User{
		Name:         nu.Name,
		Email:        nu.Email,
		PasswordHash: string(hashedPass),
	}

	// We attempt to create the new User record in the database.
	user, err := s.UserRepo.CreateUser(ctx, u)
	if err != nil {
		return models.User{}, err

	}
	return user, nil
}

// Authenticate is a method that checks a user's provided email and password against the database.
func (s *Store) Authenticate(ctx context.Context, email, password string) (jwt.RegisteredClaims,
	error) {

	// We attempt to find the User record where the email
	// matches the provided email.
	claims, err := s.UserRepo.CheckEmail(ctx, email, password)
	if err != nil {
		return jwt.RegisteredClaims{}, err
	}
	return claims, nil
}
