package api

import (
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func HandleIndex(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	WriteJSON(w, http.StatusOK, map[string]string{
		"ServiceName": "Base Microservice",
		"Version":     "1.0",
		"Mainteiner":  "PG",
	})
}

func HandleUserId(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	user := GetUserById(GetConn(), params.ByName("id"))
	WriteJSON(w, http.StatusOK, user)
}

func HandleAllusers(w http.ResponseWriter, r *http.Request, params httprouter.Params) {
	users := GetAllUsers(GetConn())
	WriteJSON(w, http.StatusOK, users)
}