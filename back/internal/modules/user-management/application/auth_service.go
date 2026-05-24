package user_mngmnt_application

import (
	"context"
	"errors"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	shared_domain_uuid "language-learning/internal/modules/shared/domain/uuid"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"

	"github.com/google/uuid"
)

var (
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)

type SignupRequest struct {
	Email       string `json:"email" required:"true"`
	Password    string `json:"password" required:"true"`
	FirstName   string `json:"first_name" required:"true"`
	LastName    string `json:"last_name" required:"true"`
	IsProfessor bool   `json:"is_professor"`
}

type SignupResponse struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
}

type LoginRequest struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}

type LoginResponse struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}

type AuthService struct {
	uuidGenerator shared_domain_uuid.UuidGenerator
	timeGenerator shared_domain_time.TimeGenerator
	users         user_mngmt_domain.Users
}

func NewAuthSrv(uuidGenerator shared_domain_uuid.UuidGenerator, timeGenerator shared_domain_time.TimeGenerator, users user_mngmt_domain.Users) *AuthService {
	return &AuthService{
		uuidGenerator: uuidGenerator,
		timeGenerator: timeGenerator,
		users:         users,
	}
}

func (as *AuthService) Login(ctx context.Context, loginRequest LoginRequest) (LoginResponse, error) {
	_, err := as.users.GetByEmail(loginRequest.Email)
	if err != nil && errors.Is(err, user_mngmt_domain.ErrUserNotFound) {
		return LoginResponse{}, ErrWrongEmailOrPassword
	}
	return LoginResponse{}, nil
}

func (as *AuthService) Signup(ctx context.Context, signupRequest SignupRequest) (SignupResponse, error) {
	_, err := as.users.GetByEmail(signupRequest.Email)
	if err == nil {
		return SignupResponse{}, user_mngmt_domain.ErrEmailAlreadyUsed
	}

	newUuid := as.uuidGenerator.Generate()
	newUser := user_mngmt_domain.NewUser(newUuid, signupRequest.Email, signupRequest.FirstName, signupRequest.FirstName, signupRequest.IsProfessor)

	err = as.users.SaveUser(newUser, signupRequest.Password)
	if err != nil {
		return SignupResponse{}, err
	}

	return SignupResponse{
		Id:    newUuid,
		Email: signupRequest.Email,
	}, nil
}
