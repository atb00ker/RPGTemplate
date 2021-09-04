package rpg

import (
	"encoding/json"
	"net/http"
	"rpgtemplate/src/controllers/config"
)

const (
	GetRPGUrl    = "/rpg"
	CreateRPGUrl = "/rpg/create"
)

func GetRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var rpgs []config.Rpgs

	config.Database.Find(&rpgs)

	rpgJson, err := json.Marshal(rpgs)
	if err != nil {
		return
	}

	response.Write(rpgJson)
}

func CreateRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(request.Body)

	var inputData createFormData
	if err := decoder.Decode(&inputData); err != nil {
		errorBody, _ := json.Marshal(responseError{Message: "Incorrect data received."})
		response.WriteHeader(http.StatusBadRequest)
		response.Write(errorBody)
		return
	}

	rpgs := config.Rpgs{RPG: inputData.RPG}
	config.Database.Create(&rpgs)
	rpgJson, err := json.Marshal(rpgs)
	if err != nil {
		errorBody, _ := json.Marshal(responseError{Message: "Creating record failed."})
		response.WriteHeader(http.StatusBadRequest)
		response.Write(errorBody)
		return
	}

	response.Write(rpgJson)
}
