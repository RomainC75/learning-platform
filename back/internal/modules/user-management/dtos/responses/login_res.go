package user_mngmt_dto_res

import "github.com/google/uuid"

type LoginResponse struct {
	Id    uuid.UUID `json:"id"`
	Email string    `json:"email"`
	Token string    `json:"token"`
}
