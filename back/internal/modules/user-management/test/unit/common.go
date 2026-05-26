package user_mngmnt_unit

import (
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	user_mngmt_dto_req "language-learning/internal/modules/user-management/dtos/requests"
	"time"

	"github.com/google/uuid"
)

var (
	newUserUuid   = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	userEmail     = "john_doe@email.com"
	userPassword  = "123456789"
	userFirstName = "John"
	userLastName  = "Doe"
	token         = "FAKE_TOKEN"
	newUserReq    = user_mngmt_dto_req.SignupRequest{
		Email:     userEmail,
		Password:  userPassword,
		FirstName: userFirstName,
		LastName:  userLastName,
		UserRole:  user_mngmt_domain.IsStudent,
	}
	savedUser = user_mngmt_domain.NewUser(newUserUuid, userEmail, userPassword, userFirstName, userLastName, user_mngmt_domain.IsStudent)
	loginReq  = user_mngmt_dto_req.LoginRequest{
		Email:    userEmail,
		Password: userPassword,
	}
	now                        = time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
	bcryptExpectedEncryptedStr = "encrypted_string"
)
