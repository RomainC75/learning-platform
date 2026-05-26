package user_mngmnt_unit

import (
	"context"
	shared_infra "language-learning/internal/modules/shared/infra"
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"testing"

	"github.com/stretchr/testify/assert"
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
// }

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
		// TODO : bcrypt password generation error
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
