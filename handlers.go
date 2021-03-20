package main

import (
	"encoding/json"
	"log"

	db "github.com/jerrydevin96/lifo-queue/database"
)

func pushHandler(data string) string {
	response := ``
	log.Println("Pushing item into queue")
	var JSONValue map[string]string
	responseJSON := make(map[string]string)
	json.Unmarshal([]byte(data), &JSONValue)
	log.Println("fetching current last index")
	index, _, err := db.GetLastRecord()
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		responseJSON["response"] = err.Error()
		responseData, _ := json.Marshal(responseJSON)
		return string(responseData)
	}
	err = db.InsertNewRecord((index+1), JSONValue["value"])
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		responseJSON["response"] = err.Error()
		responseData, _ := json.Marshal(responseJSON)
		return string(responseData)
	}
	responseJSON["response"] = "successfully pushed " + JSONValue["value"] + " into the queue"
	responseData, _ := json.Marshal(responseJSON)
	response = string(responseData)
	return response
}

func popHandler(data string) string {
	response := ``
	log.Println("Pushing item into queue")
	var JSONValue map[string]string
	responseJSON := make(map[string]string)
	json.Unmarshal([]byte(data), &JSONValue)
	log.Println("fetching current last record")
	index, value, err := db.GetLastRecord()
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		responseJSON["message"] = err.Error()
		responseJSON["value"] = "failed to fetch"
		responseData, _ := json.Marshal(responseJSON)
		return string(responseData)
	}
	log.Println("Deleting current last record")
	err = db.DeleteLastRecord(index)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		responseJSON["message"] = err.Error()
		responseJSON["value"] = "failed to fetch"
		responseData, _ := json.Marshal(responseJSON)
		return string(responseData)
	}
	responseJSON["message"] = "successfully performed POP"
	responseJSON["value"] = value
	responseData, _ := json.Marshal(responseJSON)
	response = string(responseData)
	return response
}
