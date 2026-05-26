package user_mngmt_dto_req

import user_mngmt_domain "language-learning/internal/modules/user-management/domain"

type SignupRequest struct {
	Email     string                     `json:"email" required:"true"`
	Password  string                     `json:"password" required:"true"`
	FirstName string                     `json:"first_name" required:"true"`
	LastName  string                     `json:"last_name" required:"true"`
	UserRole  user_mngmt_domain.UserRole `json:"user_role" validate:"required,numeric,min=0,max=1"`
}
