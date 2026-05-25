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
	now                = time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
	bcryptEncryptedStr = "encrypted_string"
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

func TestAuthServiceSignup(t *testing.T) {
	t.Run("should not signup if email is already used", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, false, false)

		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

		deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptEncryptedStr, false)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, users)

		_, err := authSrv.Signup(ctx, newUserReq)

		assert.EqualError(t, err, user_mngmt_domain.ErrEmailAlreadyUsed.Error())

	})

	t.Run("should return error if new user could not be saved", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, true, true)
		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

		deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptEncryptedStr, false)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, users)

		_, err := authSrv.Signup(ctx, newUserReq)

		assert.EqualError(t, err, user_mngmt_domain.ErrTryingtoSaveTheNewUser.Error())

	})
	t.Run("should signup a new user and the encrypted password saved", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, true, false)
		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)

		deterministicBcrypt := user_mngmt_infra.NewDeterministicBcrypt(bcryptEncryptedStr, false)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, deterministicBcrypt, users)

		res, err := authSrv.Signup(ctx, newUserReq)
		assert.Nil(t, err)

		savedUser := users.GetLastSavedUser()
		buffUser := user_mngmt_domain.NewUser(uuidGenerator.Generate(), newUserReq.Email, bcryptEncryptedStr, newUserReq.FirstName, newUserReq.LastName, newUserReq.IsProfessor)
		assert.Equal(t, buffUser, savedUser)

		expectedRes := user_mngmnt_application.SignupResponse{
			Id:    newUserUuid,
			Email: userEmail,
		}
		assert.Equal(t, expectedRes, res)
	})

}
