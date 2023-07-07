package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/wesley-gomes-sje/api-go-rest/database"
	"github.com/wesley-gomes-sje/api-go-rest/models"
)

func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the home page!")
}

func AllPersonalities(w http.ResponseWriter, r *http.Request) {
	var p []models.Personalitie
	database.DB.Find(&p)
	json.NewEncoder(w).Encode(p)
}

func FindPersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	var personalitie models.Personalitie
	database.DB.First(&personalitie, key)
	json.NewEncoder(w).Encode(personalitie)

}

func CreatePersonality(w http.ResponseWriter, r *http.Request) {
	var newPersonalitie models.Personalitie
	json.NewDecoder(r.Body).Decode(&newPersonalitie)
	database.DB.Create(&newPersonalitie)
	json.NewEncoder(w).Encode(newPersonalitie)
}

func DeletePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalitie models.Personalitie
	database.DB.Delete(&personalitie, id)
	fmt.Fprintf(w, "The personality with ID %v has been deleted successfully", id)
}

func EditePersonality(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	var personalitie models.Personalitie
	database.DB.First(&personalitie, id)
	json.NewDecoder(r.Body).Decode(&personalitie)
	database.DB.Save(&personalitie)
	json.NewEncoder(w).Encode(personalitie)

}
