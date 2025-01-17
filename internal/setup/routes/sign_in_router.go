package routes

import (
	"net/http"

	"github.com/Williancc1557/Oauth2.0-golang/internal/presentation/controllers"
	"github.com/Williancc1557/Oauth2.0-golang/internal/setup/adapters"
)

func SignInRouter(server *http.ServeMux) {
	signInController := &controllers.SignInController{}
	server.HandleFunc("POST /", adapters.AdaptRoute(signInController))
}
