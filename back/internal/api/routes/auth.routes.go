package routes

import (
	"language-learning/internal/api/controllers"
	"net/http"

	"github.com/gorilla/mux"
)

func AuthRoutes(mux *mux.Router) {

	authRouter := mux.PathPrefix("/auth").Subrouter()
	authCtrl := controllers.NewAuthController()

	authRouter.Handle("/signup", http.HandlerFunc(authCtrl.HandleSignup)).Methods("POST")
	authRouter.Handle("/login", http.HandlerFunc(authCtrl.HandleLogin)).Methods("POST")
}
