package user_mngmnt_unit

import (
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
	"time"

	"github.com/google/uuid"
)

var (
	newUserUuid   = uuid.MustParse("3d30a9c8-5768-4d73-ab5d-8d36cab7f03f")
	userEmail     = "john_doe@email.com"
	userPassword  = "123456789"
	userFirstName = "John"
	userLastName  = "Doe"
	newUserReq    = user_mngmnt_application.SignupRequest{
		Email:       userEmail,
		Password:    userPassword,
		FirstName:   userFirstName,
		LastName:    userLastName,
		IsProfessor: false,
	}
	savedUser = user_mngmt_domain.NewUser(newUserUuid, userEmail, userPassword, userFirstName, userLastName, false)
	loginReq  = user_mngmnt_application.LoginRequest{
		Email:    userEmail,
		Password: userPassword,
	}
	now                        = time.Date(2026, time.May, 1, 12, 0, 0, 0, time.UTC)
	bcryptExpectedEncryptedStr = "encrypted_string"
)
