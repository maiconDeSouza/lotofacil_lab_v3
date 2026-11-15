package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/model"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/service"
)

type HandlerInterface interface {
	Home(w http.ResponseWriter, r *http.Request)
	CreateNewGameFixed(w http.ResponseWriter, r *http.Request)
	UpdateGameFixed(w http.ResponseWriter, r *http.Request)
	GetGameFixed(w http.ResponseWriter, r *http.Request)
	GetAllGamesFixed(w http.ResponseWriter, r *http.Request)
	CreateNewGameFrequency(w http.ResponseWriter, r *http.Request)
	UpdateGameFrequency(w http.ResponseWriter, r *http.Request)
	GetGameFrequency(w http.ResponseWriter, r *http.Request)
	GetAllGamesFrequency(w http.ResponseWriter, r *http.Request)
	CreateNewGameRandom(w http.ResponseWriter, r *http.Request)
	UpdateGameRandom(w http.ResponseWriter, r *http.Request)
	GetGameRandom(w http.ResponseWriter, r *http.Request)
	GetAllGamesRandom(w http.ResponseWriter, r *http.Request)
	UpdateScored(w http.ResponseWriter, r *http.Request)
	ConfirmResult(w http.ResponseWriter, r *http.Request)
}

type Handler struct {
	service service.ServiceInterface
}

func NewHandler(service service.ServiceInterface) *Handler {
	return &Handler{
		service: service,
	}
}

func (h *Handler) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "ok")
}

func (h *Handler) CreateNewGameFixed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	n := model.RequestGame{}

	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.CreateNewGameFixed(n.Number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) UpdateGameFixed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	up := model.RequestUpdateGame{}

	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.UpdateGameFixed(number, up)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetGameFixed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.GetGameFixed(number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetAllGamesFixed(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gamers, errAppRes := h.service.GetAllGamesFixed()
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gamers)
}

func (h *Handler) CreateNewGameFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	n := model.RequestGame{}

	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.CreateNewGameFrequency(n.Number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) UpdateGameFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	up := model.RequestUpdateGame{}

	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.UpdateGameFrequency(number, up)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetGameFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.GetGameFrequency(number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetAllGamesFrequency(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gamers, errAppRes := h.service.GetAllGamesFrequency()
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gamers)
}

func (h *Handler) CreateNewGameRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	n := model.RequestGame{}

	if err := json.NewDecoder(r.Body).Decode(&n); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.CreateNewGameRandom(n.Number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) UpdateGameRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	up := model.RequestUpdateGame{}

	if err := json.NewDecoder(r.Body).Decode(&up); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.UpdateGameRandom(number, up)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetGameRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	g, errAppRes := h.service.GetGameRandom(number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(g)
}

func (h *Handler) GetAllGamesRandom(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	gamers, errAppRes := h.service.GetAllGamesRandom()
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(gamers)
}

func (h *Handler) UpdateScored(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	up := &model.Scored{}

	if err := json.NewDecoder(r.Body).Decode(up); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Problema na hora de processar o JSON", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	s, errAppRes := h.service.UpdateScored(up)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(s)
}

func (h *Handler) ConfirmResult(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	s := r.PathValue("number")
	number, err := strconv.Atoi(s)
	if err != nil {
		errAppRes := errapp.NewErrAppRes(err, "Parametro tem quer ser um número", http.StatusBadRequest)
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	rc, errAppRes := h.service.ConfirmResult(number)
	if errAppRes != nil {
		if errAppRes.Code == 0 {
			w.WriteHeader(http.StatusInternalServerError)
			json.NewEncoder(w).Encode(errAppRes)
			return
		}
		w.WriteHeader(errAppRes.Code)
		json.NewEncoder(w).Encode(errAppRes)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(rc)
}
