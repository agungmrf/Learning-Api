package controllers

import (
	json "encoding/json"
	"io/ioutil"
	"net/http"
	"restful-api-golang/database"
	"restful-api-golang/entity"
	"strconv"

	"github.com/gorilla/mux"
)

//GetAllPerson get all person data
func GetAllPerson(w http.ResponseWriter, r *http.Request) {
	var persons []entity.Person
	database.Connector.Find(&persons)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(persons)
	if err != nil {
		return
	}
}

//GetPersonByID returns person with specific ID
func GetPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	database.Connector.First(&person, key)
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		return
	}
}

//CreatePerson creates person
func CreatePerson(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)

	database.Connector.Create(person)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		return
	}
}

//UpdatePersonByID updates person with respective ID
func UpdatePersonByID(w http.ResponseWriter, r *http.Request) {
	requestBody, _ := ioutil.ReadAll(r.Body)
	var person entity.Person
	json.Unmarshal(requestBody, &person)
	database.Connector.Save(&person)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(person)
	if err != nil {
		return
	}
}

//DeletPersonByID delete's person with specific ID
func DeletPersonByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]

	var person entity.Person
	id, _ := strconv.ParseInt(key, 10, 64)
	database.Connector.Where("id = ?", id).Delete(&person)
	w.WriteHeader(http.StatusNoContent)
}
