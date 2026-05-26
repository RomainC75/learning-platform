package user_mngmnt_unit

import (
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"testing"

	"github.com/stretchr/testify/assert"
)

type AuthLoginTestCase struct {
	CaseMessage string
	user_mngmt_infra.InMemUsersErrors
	IsBcryptError   bool
	IsErrorExpected bool
	ExpectedError   error
	User            *user_mngmt_domain.User
}

var (
	loginTestCases = []AuthLoginTestCase{
		{
			"should not login if email is not found",
			user_mngmt_infra.InMemUsersErrors{
				ExpectedGetError: true,
			},
			false,
			true,
			user_mngmnt_application.ErrWrongEmailOrPassword,
			nil,
		},
		{
			"should not login if password is wrong found",
			user_mngmt_infra.InMemUsersErrors{
				ExpectedGetError: false,
			},
			true,
			true,
			user_mngmnt_application.ErrWrongEmailOrPassword,
			savedUser,
		},
		{
			"shoulds get the token is E/P are correct",
			user_mngmt_infra.InMemUsersErrors{
				ExpectedGetError: false,
			},
			false,
			false,
			nil,
			savedUser,
		},
	}
)

func TestAuthServiceLogin(t *testing.T) {
	for _, tc := range loginTestCases {
		t.Run(tc.CaseMessage, func(t *testing.T) {

			td := NewTestDriver(now, newUserUuid, tc.InMemUsersErrors, tc.IsBcryptError)
			if tc.User != nil {
				td.SaveUser(tc.User)
			}
			res, err := td.Login(loginReq)

			if tc.IsErrorExpected {
				assert.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				assert.Nil(t, err)

				expectedRes := user_mngmnt_application.LoginResponse{
					Id:    newUserUuid,
					Email: userEmail,
					Token: token,
				}
				assert.Equal(t, expectedRes, res)
			}

		})
	}
}

type AuthTestCase struct {
	CaseMessage string
	user_mngmt_infra.InMemUsersErrors
	IsBcryptError   bool
	IsErrorExpected bool
	ExpectedError   error
}

var (
	authTestCase = []AuthTestCase{
		{
			CaseMessage:      "should not signup if email is already used",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{false, false},
			IsBcryptError:    false,
			IsErrorExpected:  true,
			ExpectedError:    user_mngmt_domain.ErrEmailAlreadyUsed,
		},
		{
			CaseMessage:      "should return error if new user could not be saved",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{true, true},
			IsBcryptError:    false,
			IsErrorExpected:  true,
			ExpectedError:    user_mngmt_domain.ErrTryingtoSaveTheNewUser,
		},
		{
			CaseMessage:      "should signup a new user and the encrypted password saved",
			InMemUsersErrors: user_mngmt_infra.InMemUsersErrors{true, false},
			IsBcryptError:    false,
			IsErrorExpected:  false,
			ExpectedError:    user_mngmt_domain.ErrEmailAlreadyUsed,
		},
		// TODO : bcrypt password generation error
	}
)

func TestAuthServiceSignup(t *testing.T) {
	for _, tc := range authTestCase {
		t.Run(tc.CaseMessage, func(t *testing.T) {

			td := NewTestDriver(now, newUserUuid, tc.InMemUsersErrors, tc.IsBcryptError)

			res, err := td.Signup(newUserReq)

			if tc.IsErrorExpected {
				assert.EqualError(t, err, tc.ExpectedError.Error())
			} else {
				assert.Nil(t, err)

				savedUser := td.GetLastSavedUser()
				buffUser := user_mngmt_domain.NewUser(newUserUuid, newUserReq.Email, bcryptExpectedEncryptedStr, newUserReq.FirstName, newUserReq.LastName, newUserReq.UserRole)
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
