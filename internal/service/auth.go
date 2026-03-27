package service

import (
"crypto/sha256"
"encoding/hex"
"errors"

"github.com/nashirabbash/backend-pfd/internal/dto"
"github.com/nashirabbash/backend-pfd/internal/middleware"
"github.com/nashirabbash/backend-pfd/internal/model"
"github.com/nashirabbash/backend-pfd/internal/repository"
"gorm.io/gorm"
)

type AuthService struct {
	userRepo *repository.UserRepository
}

func NewAuthService(userRepo *repository.UserRepository) *AuthService {
	return &AuthService{userRepo: userRepo}
}

func (s *AuthService) Register(req *dto.RegisterRequest) (*dto.AuthResponse, error) {
	existingUser, _ := s.userRepo.FindByEmail(req.Email)
	if existingUser != nil {
		return nil, errors.New("email already registered")
	}

	hashedPassword := hashPassword(req.Password)

	user := &model.User{
		Email:    req.Email,
		Password: hashedPassword,
		Name:     req.Name,
	}

	if err := s.userRepo.Create(user); err != nil {
		return nil, err
	}

	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

func (s *AuthService) Login(req *dto.LoginRequest) (*dto.AuthResponse, error) {
	user, err := s.userRepo.FindByEmail(req.Email)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errors.New("invalid email or password")
		}
		return nil, err
	}

	if !verifyPassword(user.Password, req.Password) {
		return nil, errors.New("invalid email or password")
	}

	token, err := middleware.GenerateToken(user.ID, user.Email)
	if err != nil {
		return nil, err
	}

	return &dto.AuthResponse{
		Token: token,
		User: dto.UserResponse{
			ID:    user.ID,
			Email: user.Email,
			Name:  user.Name,
		},
	}, nil
}

func hashPassword(password string) string {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:])
}

func verifyPassword(hashedPassword, password string) bool {
	hash := sha256.Sum256([]byte(password))
	return hex.EncodeToString(hash[:]) == hashedPassword
}
