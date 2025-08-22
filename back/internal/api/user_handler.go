package api

import (
	"errors"
	"log"
	"regexp"

	"github.com/aenlemmea/garnet/back/internal/data"
	_ "golang.org/x/crypto/bcrypt"
)

type registerUserRequest struct {
	Username   string `json:"username"`
	Email      string `json:"email"`
	Password   string `json:"password"`
	Preference string `json:"preference"`
}

type UserHandler struct {
	userStore data.UserStore
	logger    *log.Logger
}

func CreateUserHandler(userStore data.UserStore, logger *log.Logger) *UserHandler {
	return &UserHandler{
		userStore: userStore,
		logger:    logger,
	}
}

// Frontend does have validation, here the backend ensures that data layer is not affected by slip-ups.
func (uh *UserHandler) validateRegisterRequest(reg *registerUserRequest) error {
	if reg.Username == "" {
		return errors.New("Username is required")
	}

	if len(reg.Username) > 50 {
		return errors.New("Username cannot be greater than 50 characters")
	}

	if reg.Email == "" {
		return errors.New("Email is mandatory")
	}

	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(reg.Email) {
		return errors.New("Invalid email format")
	}

	return nil
}
