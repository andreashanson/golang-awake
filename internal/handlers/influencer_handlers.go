package handlers

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/andreashanson/golang-awake/internal/influencers"
	"github.com/go-chi/chi"
)

type Influencer interface {
	GetAll() ([]influencers.Influencer, error)
	GetByID(id string) (influencers.Influencer, error)
	Create(name, lastname, email string) error
}

type InfluencerHandler struct {
	Repo Influencer
}

func NewInfluencerHandler(ir Influencer) *InfluencerHandler {
	return &InfluencerHandler{Repo: ir}
}

func (ih *InfluencerHandler) GetAll(w http.ResponseWriter, r *http.Request) {
	allInfluencers, _ := ih.Repo.GetAll()
	if err := json.NewEncoder(w).Encode(&allInfluencers); err != nil {
		log.Println("Failed to Encode json influencers in GetAll:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}

func (ih InfluencerHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	influencer, err := ih.Repo.GetByID(id)
	if err != nil {
		log.Println("Failed to GetByID:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(&influencer); err != nil {
		log.Println("Failed to Encode json influencer in GetByID:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func (ih InfluencerHandler) Create(w http.ResponseWriter, r *http.Request) {
	var i influencers.Influencer
	if err := json.NewDecoder(r.Body).Decode(&i); err != nil {
		log.Println("Failed to Decode json influencer in create:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	if err := ih.Repo.Create(i.Name, i.Lastname, i.Email); err != nil {
		log.Println("Failed to create new influencer:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
