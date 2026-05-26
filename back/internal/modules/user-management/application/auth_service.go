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
	bcryptCost              = 10
)

type SignupRequest struct {
	Email     string                     `json:"email" required:"true"`
	Password  string                     `json:"password" required:"true"`
	FirstName string                     `json:"first_name" required:"true"`
	LastName  string                     `json:"last_name" required:"true"`
	UserRole  user_mngmt_domain.UserRole `json:"user_role" validate:"required,numeric,min=0,max=1"`
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
	bcrypt        user_mngmt_domain.Bcrypt
	jwt           user_mngmt_domain.Jwt
	users         user_mngmt_domain.Users
}

func NewAuthSrv(uuidGenerator shared_domain_uuid.UuidGenerator, timeGenerator shared_domain_time.TimeGenerator, bcrypt user_mngmt_domain.Bcrypt, jwt user_mngmt_domain.Jwt, users user_mngmt_domain.Users) *AuthService {
	return &AuthService{
		uuidGenerator: uuidGenerator,
		timeGenerator: timeGenerator,
		bcrypt:        bcrypt,
		jwt:           jwt,
		users:         users,
	}
}

func (as *AuthService) Login(ctx context.Context, loginRequest LoginRequest) (LoginResponse, error) {
	foundUser, err := as.users.GetByEmail(loginRequest.Email)
	if err != nil && errors.Is(err, user_mngmt_domain.ErrUserNotFound) {
		return LoginResponse{}, ErrWrongEmailOrPassword
	}
	err = foundUser.IsPasswordValid(as.bcrypt.CompareHashAndPassword, loginRequest.Password)
	if err != nil {
		return LoginResponse{}, ErrWrongEmailOrPassword
	}
	token, err := as.jwt.CreateToken(foundUser.Email())
	if err != nil {
		return LoginResponse{}, err
	}
	return LoginResponse{
		Id:    foundUser.Id(),
		Email: foundUser.Email(),
		Token: token,
	}, nil
}

func (as *AuthService) Signup(ctx context.Context, signupRequest SignupRequest) (SignupResponse, error) {
	_, err := as.users.GetByEmail(signupRequest.Email)
	if err == nil {
		return SignupResponse{}, user_mngmt_domain.ErrEmailAlreadyUsed
	}

	encryptedPassword, err := as.bcrypt.GenerateFromPassword([]byte(signupRequest.Password), bcryptCost)
	if err != nil {
		return SignupResponse{}, user_mngmt_domain.ErrTryingToGenerateBcryptPassword
	}

	newUuid := as.uuidGenerator.Generate()
	newUser := user_mngmt_domain.NewUser(newUuid, signupRequest.Email, string(encryptedPassword), signupRequest.FirstName, signupRequest.LastName, signupRequest.UserRole)

	err = as.users.Save(newUser)
	if err != nil {
		return SignupResponse{}, err
	}

	return SignupResponse{
		Id:    newUuid,
		Email: signupRequest.Email,
	}, nil
}
