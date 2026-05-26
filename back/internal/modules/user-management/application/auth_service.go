package user_mngmnt_application

import (
	"context"
	"errors"
	shared_domain_time "language-learning/internal/modules/shared/domain/time"
	shared_domain_uuid "language-learning/internal/modules/shared/domain/uuid"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_dto_req "language-learning/internal/modules/user-management/dtos/requests"
	user_mngmt_dto_res "language-learning/internal/modules/user-management/dtos/responses"
)

var (
	ErrWrongEmailOrPassword = errors.New("wrong email or password")
)

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

func (as *AuthService) Login(ctx context.Context, loginRequest user_mngmt_dto_req.LoginRequest) (user_mngmt_dto_res.LoginResponse, error) {
	foundUser, err := as.users.GetByEmail(loginRequest.Email)
	if err != nil && errors.Is(err, user_mngmt_domain.ErrUserNotFound) {
		return user_mngmt_dto_res.LoginResponse{}, ErrWrongEmailOrPassword
	}
	err = foundUser.IsPasswordValid(as.bcrypt.CompareHashAndPassword, loginRequest.Password)
	if err != nil {
		return user_mngmt_dto_res.LoginResponse{}, ErrWrongEmailOrPassword
	}
	token, err := as.jwt.CreateToken(foundUser.Email(), foundUser.Role())
	if err != nil {
		return user_mngmt_dto_res.LoginResponse{}, err
	}
	return user_mngmt_dto_res.LoginResponse{
		Id:    foundUser.Id(),
		Email: foundUser.Email(),
		Token: token,
	}, nil
}

func (as *AuthService) Signup(ctx context.Context, signupRequest user_mngmt_dto_req.SignupRequest) (user_mngmt_dto_res.SignupResponse, error) {
	_, err := as.users.GetByEmail(signupRequest.Email)
	if err == nil {
		return user_mngmt_dto_res.SignupResponse{}, user_mngmt_domain.ErrEmailAlreadyUsed
	}

	encryptedPassword, err := as.bcrypt.GenerateFromPassword([]byte(signupRequest.Password), user_mngmt_domain.BcryptCost)
	if err != nil {
		return user_mngmt_dto_res.SignupResponse{}, user_mngmt_domain.ErrTryingToGenerateBcryptPassword
	}

	newUuid := as.uuidGenerator.Generate()
	newUser := user_mngmt_domain.NewUser(newUuid, signupRequest.Email, string(encryptedPassword), signupRequest.FirstName, signupRequest.LastName, signupRequest.UserRole)

	err = as.users.Save(newUser)
	if err != nil {
		return user_mngmt_dto_res.SignupResponse{}, err
	}

	return user_mngmt_dto_res.SignupResponse{
		Id:    newUuid,
		Email: signupRequest.Email,
	}, nil
}
