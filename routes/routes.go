package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dcao96/Simple-Go-Key-Server/parsing"
	"github.com/dcao96/Simple-Go-Key-Server/types"
)

func GetHandler(w http.ResponseWriter, r *http.Request) {
	data := r.Header.Get("data")
	bytes := []byte(data)
	var get types.Get
	json.Unmarshal(bytes, &get)

	key, err := parsing.GetPublicKey(get.FacebookID)

	var response types.GetPublicKeyResponse
	var errorResponse types.ErrorResponse
	if err != nil {
		fmt.Printf("Error getting user: %v\n", err)
		errorResponse = types.ErrorResponse{true, err.Error()}
	} else {
		fmt.Printf("Getting user: %v\n", get.FacebookID)
		errorResponse = types.ErrorResponse{false, ""}
	}

	response = types.GetPublicKeyResponse{&errorResponse, key}

	b, err := json.Marshal(response)
	fmt.Fprintf(w, "%s", b)
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()
	fmt.Println(r.Form())
	fmt.Println(r.FormValue("data"))

	data := r.Header.Get("data")
	bytes := []byte(data)
	var user types.User
	json.Unmarshal(bytes, &user)
	// userData := f.(map[string]interface{})
	// fmt.Println("data: ", userData)

	err := parsing.AddUser(user.FacebookID, user.PublicKey)

	var errorResponse types.ErrorResponse
	if err != nil {
		fmt.Printf("Error adding user: %v\n", err)
		errorResponse = types.ErrorResponse{true, err.Error()}
	} else {
		fmt.Printf("Added user: %v\n", user.FacebookID)
		errorResponse = types.ErrorResponse{false, ""}
	}

	//write json to return here
	b, err := json.Marshal(errorResponse)
	fmt.Fprintf(w, "%s", b)
}
