package user_mngmt_dto_req

type LoginRequest struct {
	Email    string `json:"email" required:"true"`
	Password string `json:"password" required:"true"`
}
