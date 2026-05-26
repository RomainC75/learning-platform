package user_mngmt_domain

import "errors"

var (
	ErrCreatingToken = errors.New("could not create token")
)

type Jwt interface {
	CreateToken(username string) (string, error)
}
