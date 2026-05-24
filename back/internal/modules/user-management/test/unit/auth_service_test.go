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
	newUserEmail = "john_doe@email.com"
	newUserReq   = user_mngmnt_application.SignupRequest{
		Email:       newUserEmail,
		Password:    "123456789",
		FirstName:   "John",
		LastName:    "Doe",
		IsProfessor: false,
	}
	now = time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
)

func TestAuthService(t *testing.T) {
	t.Run("should not signup if email is already used", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, true, false)
		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, users)

		_, err := authSrv.Signup(ctx, newUserReq)

		assert.EqualError(t, err, user_mngmt_domain.ErrEmailAlreadyUsed)

	})

	t.Run("should return error if new user could not be saved", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, false, true)
		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, users)

		_, err := authSrv.Signup(ctx, newUserReq)

		assert.EqualError(t, err, user_mngmt_domain.ErrTryingtoSaveTheNewUser)

	})
	t.Run("should signup a new user", func(t *testing.T) {

		users := user_mngmt_infra.NewInMemUsers(nil, false, false)
		ctx := context.Background()

		timeGenerator := shared_infra.NewDeterministicTimeGenerator(now)
		uuidGenerator := shared_infra.NewInMemUuidGenerator(newUserUuid)
		authSrv := user_mngmnt_application.NewAuthSrv(uuidGenerator, timeGenerator, users)

		res, err := authSrv.Signup(ctx, newUserReq)

		assert.Nil(t, err)

		expectedRes := user_mngmnt_application.SignupResponse{
			Id:    newUserUuid,
			Email: newUserEmail,
		}
		assert.Equal(t, expectedRes, res)
	})
}
