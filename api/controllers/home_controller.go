package controllers

import (
	"net/http"

	"github.com/Muadhim/app-todo-list/api/responses"
)

func (server *Server) Home(w http.ResponseWriter, r *http.Request) {
	responses.JSON(w, http.StatusOK, "Wolcome To This Awesome API")
}