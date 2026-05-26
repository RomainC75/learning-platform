package user_mngmt_infra

import (
	user_mngmt_domain "language-learning/internal/modules/user-management/domain"
)

type DeterministicBcrypt struct {
	encryptedStr  string
	expectedError bool
	// lastClearPassword []byte
}

func NewDeterministicBcrypt(encryptedStr string, expectedError bool) *DeterministicBcrypt {
	return &DeterministicBcrypt{
		encryptedStr:  encryptedStr,
		expectedError: expectedError,
	}
}

func (db *DeterministicBcrypt) CompareHashAndPassword(hashedPassword, password []byte) error {
	if db.expectedError {
		return user_mngmt_domain.ErrInvalidPass
	}
	return nil
}

func (db *DeterministicBcrypt) GenerateFromPassword(password []byte, cost int) ([]byte, error) {
	if db.expectedError {
		// db.lastClearPassword = password
		return []byte{}, user_mngmt_domain.ErrTryingToGenerateBcryptPassword
	}
	return []byte(db.encryptedStr), nil
}
