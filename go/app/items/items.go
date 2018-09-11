package items

import (
	"encoding/json"
	"net/http"
)

func List(response http.ResponseWriter, request *http.Request) {
	response.Header().Set("Content-Type", "application/json")

	newItem := Item{}
	newItem.Title = "Tester"
	newItem.Icon = "file"

	dataSet := []Item{}
	dataSet = append(dataSet, newItem)

	json.NewEncoder(response).Encode(dataSet)
}
