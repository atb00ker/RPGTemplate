package rpg

import (
	"encoding/json"
	"net/http"
	"rpgtemplate/src/controllers/config"

	"github.com/gorilla/mux"
)

const (
	GetRPGUrl     = "/rpg"
	CreateRPGUrl  = "/rpg/create"
	DetailsRPGUrl = "/rpg/{rpgId}"
)

func GetRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	var rpgs []config.Rpgs

	config.Database.Find(&rpgs)

	rpgJson, err := json.Marshal(rpgs)
	if err != nil {
		handleError(err, "Could not fetch data, contact site administrator.", response)
	}

	response.Write(rpgJson)
}

func CreateRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	decoder := json.NewDecoder(request.Body)

	var inputData createFormData
	if err := decoder.Decode(&inputData); err != nil {
		handleError(err, "Incorrect data received.", response)
	}

	rpgData := config.Rpgs{RPG: inputData.RPG}
	result := config.Database.Create(&rpgData)
	if result.Error != nil {
		handleError(result.Error, "Creating record failed.", response)
	}
	rpgJson, _ := json.Marshal(rpgData)
	response.Write(rpgJson)
}

func UpdateRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	decoder := json.NewDecoder(request.Body)

	var inputData createFormData
	if err := decoder.Decode(&inputData); err != nil {
		handleError(err, "Incorrect data received.", response)
	}
	result := config.Database.Model(&config.Rpgs{}).Where("id = ?", vars["rpgId"]).Update("rpg", inputData.RPG)
	if result.Error != nil {
		handleError(result.Error, "Updating record failed.", response)
	}
	rpgJson, _ := json.Marshal(inputData)
	response.Write(rpgJson)
}

func DeleteRPG(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(request)
	result := config.Database.Delete(&config.Rpgs{}, vars["rpgId"])
	if result.Error != nil {
		handleError(result.Error, "Deleting record failed.", response)
	}

	rpgJson, _ := json.Marshal("Done")
	response.Write(rpgJson)
}
