package main

import (
	"election_system/entity"
	"election_system/service"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strings"
)

var ser = service.Connection{}
var finalResponse entity.Response

func init() {
	ser.Server = "mongodb://localhost:27017"
	ser.Database = "ElectionSystem"
	ser.Collection1 = "UserDetails"

	ser.Connect()
}
func generateToken(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var loginDetails entity.LoginDetails
	if err := json.NewDecoder(r.Body).Decode(&loginDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	}

	if result, err := ser.GenerateToken(loginDetails); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func saveUserDetails(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "POST" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody entity.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.SaveUserDetails(dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func verifyUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	var dataBody entity.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateUserDetails(dataBody, id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func updateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	email := segments[len(segments)-1]

	var dataBody1 entity.User
	if err := json.NewDecoder(r.Body).Decode(&dataBody1); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UpdateUser(dataBody1, email); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func searchUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody2 entity.UserSearch
	if err := json.NewDecoder(r.Body).Decode(&dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UseSearch(dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func multipleUserSearch(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "GET" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody2 entity.MultipleUserSearch
	if err := json.NewDecoder(r.Body).Decode(&dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.MultiSearch(dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "DELETE" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}

	var dataBody2 entity.DeleteUser
	if err := json.NewDecoder(r.Body).Decode(&dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid Request")
		return
	}

	if result, err := ser.UserDelete(dataBody2); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func deactivateUser(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()

	if r.Method != "PUT" {
		respondWithError(w, http.StatusBadRequest, "Invalid Method")
		return
	}
	path := r.URL.Path
	segments := strings.Split(path, "/")
	id := segments[len(segments)-1]

	if result, err := ser.UserDeactivate(id); err != nil {
		respondWithError(w, http.StatusBadRequest, fmt.Sprintf("%v", err))
	} else {
		respondWithJson(w, http.StatusAccepted, result, "")
	}
}

func main() {
	http.HandleFunc("/generate-token", generateToken)
	http.HandleFunc("/user/save-user", saveUserDetails)
	http.HandleFunc("/user/verifyUser/", verifyUser)
	http.HandleFunc("/user/updateUser/", updateUser)
	http.HandleFunc("/user/searchUser", searchUser)
	http.HandleFunc("/user/search-multiple-user", multipleUserSearch)
	http.HandleFunc("/delete-user", deleteUser)
	http.HandleFunc("/user/deactivate-user/", deactivateUser)
	log.Println("Server started at 8080")
	http.ListenAndServe(":8080", nil)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg}, "error")
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}, err string) {
	if err == "error" {
		finalResponse.Success = "false"
	} else {
		finalResponse.Success = "true"
	}
	finalResponse.SucessCode = fmt.Sprintf("%v", code)
	finalResponse.Response = payload
	response, _ := json.Marshal(finalResponse)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}
