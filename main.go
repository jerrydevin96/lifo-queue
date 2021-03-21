package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/jerrydevin96/lifo-queue/startup"
)

func main() {
	log.Println("Starting the Application")
	err := startup.AppStartup()
	if err != nil {
		log.Panic("[APPLICATION STARTUP ERROR] " + err.Error())
	}
	log.Println("App startup completed starting API service")
	http.HandleFunc("/v1/push", push)
	http.HandleFunc("/v1/pop", pop)
	log.Println("Starting API service on port 8080")
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Panic("[API STARTUP ERROR] " + err.Error())
	}
}

func push(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	JSONVal, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Println("[ERROR] " + err.Error())
		fmt.Fprint(w, `{response: `+err.Error()+`}`)
	}
	response := pushHandler(string(JSONVal))
	fmt.Fprint(w, response)
}

func pop(w http.ResponseWriter, r *http.Request) {
	(w).Header().Set("Access-Control-Allow-Origin", "*")
	response := popHandler()
	fmt.Fprint(w, response)
}
