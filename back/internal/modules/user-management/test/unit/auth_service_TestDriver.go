package user_mngmnt_unit

import (
	"context"
	shared_infra "language-learning/internal/modules/shared/infra"
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"time"

	"github.com/google/uuid"
)

type AuthTestDriver struct {
	users   *user_mngmt_infra.InMemUsers
	authSrv *user_mngmnt_application.AuthService
}

func NewTestDriver(now time.Time, newUserUuid uuid.UUID, imUserE user_mngmt_infra.InMemUsersErrors, isBcryptError bool) *AuthTestDriver {
	users := user_mngmt_infra.NewInMemUsers(imUserE)

	timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
	uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

	deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptExpectedEncryptedStr, isBcryptError)
	inMemJwt := user_mngmt_infra.NewInMemJWT(token)
	authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, inMemJwt, users)

	return &AuthTestDriver{
		users:   users,
		authSrv: authSrv,
	}
}

func (atd *AuthTestDriver) Signup(signupReq user_mngmnt_application.SignupRequest) (user_mngmnt_application.SignupResponse, error) {
	ctx := context.Background()
	return atd.authSrv.Signup(ctx, signupReq)
}

func (atd *AuthTestDriver) Login(loginReq user_mngmnt_application.LoginRequest) (user_mngmnt_application.LoginResponse, error) {
	ctx := context.Background()
	return atd.authSrv.Login(ctx, loginReq)
}

func (atd *AuthTestDriver) GetLastSavedUser() *user_mngmt_domain.User {
	return atd.users.GetLastSavedUser()
}

func (atd *AuthTestDriver) SaveUser(user *user_mngmt_domain.User) {
	atd.users.Save(user)
}
