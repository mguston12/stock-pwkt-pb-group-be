package stock

import (
	"context"
	"stock/internal/entity/stock"
	"stock/pkg/auth"
	"stock/pkg/errors"

	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", errors.Wrap(err, "[SERVICE][HashPassword]")
	}
	return string(hashedPassword), nil
}

func (s Service) GetUserByUsername(ctx context.Context, username string) (stock.User, error) {
	user, err := s.data.GetUserByUsername(ctx, username)
	if err != nil {
		return user, errors.Wrap(err, "[SERVICE][GetUserByUsername]")
	}

	return user, nil
}

// func (s Service) MatchPassword(ctx context.Context, login stock.User) error {
// 	user, err := s.GetUserByUsername(ctx, login.Username)
// 	if err != nil {
// 		return errors.New("Username tidak ditemukan.")
// 	}

// 	if user.Password == "p4ssw0rd" {
// 		return errors.New("Buat Password dulu")
// 	}

// 	// Compare the hashed password stored in DB with the provided one
// 	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(login.Password))
// 	if err != nil {
// 		return errors.New("Salah Password")
// 	}

// 	return nil
// }

func (s Service) CreateUser(ctx context.Context, user stock.User) error {
	// Hash the password before saving
	hashedPassword, err := HashPassword(user.Password)
	if err != nil {
		return err
	}

	user.Password = hashedPassword

	// Store the user with hashed password
	err = s.data.CreateUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][CreateUser]")
	}

	return nil
}

func (s Service) UpdateUser(ctx context.Context, user stock.User) error {
	// If the password has been updated, hash it before updating the user
	if user.Password != "" {
		hashedPassword, err := HashPassword(user.Password)
		if err != nil {
			return err
		}
		user.Password = hashedPassword
	}

	err := s.data.UpdateUser(ctx, user)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][UpdateUser]")
	}

	return nil
}

func (s Service) DeleteUser(ctx context.Context, username string) error {
	err := s.data.DeleteUser(ctx, username)
	if err != nil {
		return errors.Wrap(err, "[SERVICE][DeleteUser]")
	}

	return nil
}

func (s Service) Login(ctx context.Context, username, password string) (string, error) {
	user, err := s.GetUserByUsername(ctx, username)
	if err != nil {
		return "", errors.New("Username tidak ditemukan.")
	}

	if user.Password == "p4ssw0rd" {
		return "", errors.New("Buat Password dulu")
	}

	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil {
		return "", errors.New("Salah Password")
	}

	token, err := auth.GenerateJWT(user.Username, user.Role) // Role must be stored in DB
	if err != nil {
		return "", errors.Wrap(err, "[SERVICE][Login][JWT]")
	}

	return token, nil
}
