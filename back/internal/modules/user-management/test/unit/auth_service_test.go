package user_mngmnt_unit

import (
	"context"
	shared_infra "language-learning/internal/modules/shared/infra"
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
)

var (
	newUserUuid  = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	userEmail    = "john_doe@email.com"
	userPassword = "123456789"
	newUserReq   = user_mngmnt_application.SignupRequest{
		Email:       userEmail,
		Password:    userPassword,
		FirstName:   "John",
		LastName:    "Doe",
		IsProfessor: false,
	}
	loginReq = user_mngmnt_application.LoginRequest{
		Email:    userEmail,
		Password: userPassword,
	}
	now                        = time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
	bcryptExpectedEncryptedStr = "encrypted_string"
)

// func TestAuthServiceLogin(t *testing.T) {
// 	t.Run("should not Login if email is wrong", func(t *testing.T) {
// 		users := user_mngmt_infra.NewInMemUsers(nil, true, false)
// 		ctx := context.Background()

// 		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
// 		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)
// 		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, users)

// 		_, err := authSrv.Login(ctx, loginReq)
// 		assert.EqualError(t, err, user_mngmnt_application.ErrWrongEmailOrPassword.Error())
// 	})
// 	t.Run("should not Login if password is wrong", func(t *testing.T) {
// 		users := user_mngmt_infra.NewInMemUsers(nil, false, false)
// 		ctx := context.Background()

// 		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
// 		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)
// 		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, users)

// 		_, err := authSrv.Login(ctx, loginReq)
// 		assert.EqualError(t, err, user_mngmnt_application.ErrWrongEmailOrPassword.Error())
// 	})
// }

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

type AuthTestCase struct {
	CaseMessage string
	user_mngmt_infra.InMemUsersErrors
	BcryptError     bool
	IsErrorExpected bool
	ExpectedError   error
}

var (
	authTestCase = []AuthTestCase{
		{
			CaseMessage:      "should not signup if email is already used",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{false, false},
			BcryptError:      false,
			IsErrorExpected:  true,
			ExpectedError:    user_mngmt_domain.ErrEmailAlreadyUsed,
		},
		{
			CaseMessage:      "should return error if new user could not be saved",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{true, true},
			BcryptError:      false,
			IsErrorExpected:  true,
			ExpectedError:    user_mngmt_domain.ErrTryingtoSaveTheNewUser,
		},
		{
			CaseMessage:      "should signup a new user and the encrypted password saved",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{true, false},
			BcryptError:      false,
			IsErrorExpected:  false,
			ExpectedError:    user_mngmt_domain.ErrEmailAlreadyUsed,
		},
	}
)

func TestAuthServiceSignup(t *testing.T) {
	for _, tc := range authTestCase {
		t.Run(tc.CaseMessage, func(t *testing.T) {

			users := user_mngmt_infra.NewInMemUsers(tc.InMemUsersErrors)
			ctx := context.Background()

			timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
			uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

			deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptExpectedEncryptedStr, tc.BcryptError)
			authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, users)

			res, err := authSrv.Signup(ctx, newUserReq)

			if tc.IsErrorExpected {
				assert.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				assert.Nil(t, err)

				savedUser := users.GetLastSavedUser()
				buffUser := user_mngmt_domain.NewUser(uuidGenerator.Generate(), newUserReq.Email, bcryptExpectedEncryptedStr, newUserReq.FirstName, newUserReq.LastName, newUserReq.IsProfessor)
				assert.Equal(t, buffUser, savedUser)

				expectedRes := user_mngmnt_application.SignupResponse{
					Id:    newUserUuid,
					Email: userEmail,
				}
				assert.Equal(t, expectedRes, res)
			}

		})
	}

}
