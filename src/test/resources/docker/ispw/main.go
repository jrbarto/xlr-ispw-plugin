package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
)

// Release struct used to return json after createRelease is called
type Release struct {
	ReleaseID string `json:"releaseId"`
	URL       string `json:"url"`
}

// ReleaseInformation struct used to return json after getReleaseInformation is called
type ReleaseInformation struct {
	RelOutputId   string `json:"releaseId"`
	Application   string `json:"application"`
	Stream        string `json:"stream"`
	Description   string `json:"description"`
	Owner         string `json:"owner"`
	WorkRefNumber string `json:"workRefNumber"`
}

// SetInformation struct used to return json after getSetInformation is called
type SetInformation struct {
	SetOutputId              string `json:"setid"`
	Application              string `json:"applicationId"`
	Stream                   string `json:"streamName"`
	Description              string `json:"description"`
	Owner                    string `json:"owner"`
	StartDate                string `json:"startDate"`
	StartTime                string `json:"startTime"`
	DeployActivationDate     string `json:"deployActiveDate"`
	DeployActivationTime     string `json:"deployActiveTime"`
	DeployImplementationDate string `json:"deployImplementationDate"`
	DeployImplementationTime string `json:"deployImplementationTime"`
	State                    string `json:"state"`
}

// RegressResponse struct used to retun json after regress is called
type RegressResponse struct {
	SetID   string `json:"setId"`
	Message string `json:"message"`
	URL     string `json:"url"`
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/ispw/ispw/releases/", CreateRelease).Methods("POST")
	router.HandleFunc("/ispw/ispw/releases/{release_id}", GetReleaseInformation).Methods("GET")
	//router.HandleFunc("/ispw/ispw/releases/{release_id}/tasks/promote?level={level}", Promote).Methods("POST")
	//router.HandleFunc("/ispw/ispw/releases/{release_id}/tasks/deploy?level={level}", Deploy).Methods("POST")
	router.HandleFunc("/ispw/ispw/sets/{set_id}", GetSetInformation).Methods("GET")
	router.HandleFunc("/ispw/ispw/releases/relid/tasks/regress", Regress).Methods("POST")

	log.Fatal(http.ListenAndServe(":8080", router))
}

// CreateRelease sends a dummy response back
func CreateRelease(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	c := Release{"relid", "http://ispw:8080/ispw/ispw/releases/relid"}
	outgoingJSON, err := json.Marshal(c)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))
}

// GetReleaseInformation sends a dummy response back
func GetReleaseInformation(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	c := ReleaseInformation{"relid", "app", "stream", "something", "xebia", "1234"}
	outgoingJSON, err := json.Marshal(c)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))
}

// GetSetInformation sends a dummy response back
func GetSetInformation(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	c := SetInformation{"someId","app","stream","something","xebia","08/10", "11am","09/10","11am","10/10","11am", "DONE"}
	outgoingJSON, err := json.Marshal(c)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))
}

// Regress sends a dummy response back
func Regress(res http.ResponseWriter, req *http.Request) {
	res.Header().Set("Content-Type", "application/json")
	c := RegressResponse{"ISPW2345", "This worked", "http://foobarsoft.com/ispw/w3t/sets/s0123456"}
	outgoingJSON, err := json.Marshal(c)
	if err != nil {
		log.Println(err.Error())
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	res.WriteHeader(http.StatusCreated)
	fmt.Fprint(res, string(outgoingJSON))

}
