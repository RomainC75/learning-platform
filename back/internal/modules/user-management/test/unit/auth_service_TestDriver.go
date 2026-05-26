package user_mngmnt_unit

import (
	"context"
	shared_infra "language-learning/internal/modules/shared/infra"
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"time"

	"github.com/google/uuid"
)

type AuthTestDriver struct {
	users   *user_mngmt_infra.InMemUsers
	authSrv *user_mngmnt_application.AuthService
}

func NewTestDriver(now time.Time, newUserUuid uuid.UUID) *AuthTestDriver {
	users := user_mngmt_infra.NewInMemUsers(user_mngmt_infra.InMemUsersErrors{})

	timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
	uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

	deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptExpectedEncryptedStr, false)
	authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, users)

	return &AuthTestDriver{
		users:   users,
		authSrv: authSrv,
	}
}

func (atd *AuthTestDriver) Signup(signupReq user_mngmnt_application.SignupRequest) (user_mngmnt_application.SignupResponse, error) {
	ctx := context.Background()
	return atd.authSrv.Signup(ctx, signupReq)
}
