package controllers

import (
	"context"
	"encoding/json"
	api_utils "language-learning/internal/api/utils"
	validatorHandler "language-learning/internal/api/validator"
	shared_infra "language-learning/internal/modules/shared/infra"
	user_mngmnt_application "language-learning/internal/modules/user-management/application"
	user_mngmt_dto_req "language-learning/internal/modules/user-management/dtos/requests"
	user_mngmt_infra "language-learning/internal/modules/user-management/infra"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type AuthCtrl struct {
	authSrv   user_mngmnt_application.AuthService
	validator *validator.Validate
}

func NewAuthController() *AuthCtrl {
	return &AuthCtrl{
		authSrv: *user_mngmnt_application.NewAuthSrv(
			shared_infra.NewUuidGenerator(),
			shared_infra.NewTimeGenerator(),
			user_mngmt_infra.NewBcrypt(),
			user_mngmt_infra.NewJWT(),
			user_mngmt_infra.NewInMemUsers(),
		),
		validator: validatorHandler.GetValidator(),
	}
}

func (ac *AuthCtrl) HandleLogin(w http.ResponseWriter, r *http.Request) {
	var req user_mngmt_dto_req.LoginRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ac.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	logRes, err := ac.authSrv.Login(ctx, req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	api_utils.JsonResponse(w, logRes)
}

func (ac *AuthCtrl) HandleSignup(w http.ResponseWriter, r *http.Request) {
	var req user_mngmt_dto_req.SignupRequest
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = ac.validator.Struct(req)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	ctx := context.Background()
	logRes, err := ac.authSrv.Signup(ctx, req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	api_utils.JsonResponse(w, logRes)
}
