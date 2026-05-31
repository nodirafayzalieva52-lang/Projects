package handlers

import (
	"encoding/json"
	"errors"
	"net/http"
	"project/logger"
	"project/models"
	"project/storage"
	"strconv"

	"go.uber.org/zap"
)

type UserHandler struct {
	storage *storage.UserStorage
}

var internalError = errors.New("internal error")
var badRequest = errors.New("bad request")

func NewUserHandler(storage *storage.UserStorage) *UserHandler{
	return &UserHandler{
		storage: storage,
	}
}

func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.storage.GetAll()
	if err != nil {
		logger.L.Error("h.storage.GetAll:", zap.Error(err))
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)

	err = json.NewEncoder(w).Encode(users)
	if err != nil {
		logger.L.Error("json.Encode:", zap.Error(err))
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
		return
	}
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil{
		http.Error(w,badRequest.Error(),http.StatusBadRequest)
		return
	}

	user, err := h.storage.GetByID(id)
	if err != nil {
		logger.L.Error("h.storage.GetByID:", zap.Error(err))
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		logger.L.Error("json.Encode:", zap.Error(err))
	}
}

func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user models.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,badRequest.Error(),http.StatusBadRequest)
		return
	}

	err = h.storage.Create(user)
	if err != nil {
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(models.Response{
		Message: "user created",
	},
	)
	if err != nil {
		logger.L.Error("json.Encode:", zap.Error(err))
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
	}
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idStr := r.PathValue("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w,badRequest.Error(),http.StatusBadRequest)
		return
	}	

	var user models.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w,badRequest.Error(),http.StatusBadRequest)
		return
	}

	err = h.storage.Update(id,user)
	if err != nil {
		http.Error(w,internalError.Error(),http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}
