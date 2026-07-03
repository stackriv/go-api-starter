package controller

import (
	"log"
	"net/http"
	"os"

	"github.com/stackriv/go-api-starter/internal/business/service"
	"github.com/stackriv/go-api-starter/internal/pkg"
)

type UserController struct {
	Service service.UserService
}

func (u *UserController) Register(w http.ResponseWriter, r *http.Request) {
	err := pkg.Env()
	if err != nil {
		log.Println("Error loading environment variables:", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}
}

func (u *UserController) UserRoutes(routes *http.ServeMux) *http.ServeMux {
	err := pkg.Env()
	if err != nil {
		log.Println("Error loading environment variables:", err)
		return routes
	}

	routes.HandleFunc(os.Getenv("DEFAULT_API_URI")+"/register", u.Register)

	return routes
}
