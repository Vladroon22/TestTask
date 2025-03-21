package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gthub.com/Vladroon22/TestTask/internal/entity"
	"gthub.com/Vladroon22/TestTask/internal/service"
	"gthub.com/Vladroon22/TestTask/internal/utils"
)

type Handler struct {
	srv service.Servicer
}

func NewHandler(s service.Servicer) *Handler {
	return &Handler{
		srv: s,
	}
}

func (h *Handler) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	if ok := utils.ValidateEmail(user.Email); !ok {
		http.Error(w, "wrong format email", http.StatusBadRequest)
		log.Println("wrong format email")
		return
	}

	if ok := utils.ValidatePhone(user.Phone); !ok {
		http.Error(w, "wrong format phone", http.StatusBadRequest)
		log.Println("wrong format phone")
		return
	}

	if err := h.srv.CreateUser(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writeJSON(w, http.StatusOK, "User created successfully")
}

func (h *Handler) GetAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])

	user, err := h.srv.GetUser(r.Context(), ID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}

	writeJSON(w, http.StatusOK, user)
}

func (h *Handler) UpdateAccount(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	ID, _ := strconv.Atoi(vars["id"])

	var user entity.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		log.Println(err)
		return
	}
	user.ID = ID
	if _, err := h.srv.UpdateUser(r.Context(), user); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		log.Println(err)
		return
	}

	writeJSON(w, http.StatusOK, fmt.Sprintf("user with %d id is updated", ID))
}

func writeJSON(w http.ResponseWriter, status int, a interface{}) error {
	w.Header().Add("Content-Type", "application/json")
	w.WriteHeader(status)

	return json.NewEncoder(w).Encode(a)
}
