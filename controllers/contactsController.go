package controllers

import (
	"net/http"
	"build-rest-api/models"
	"encoding/json"
	"github.com/gorilla/mux"
	"strconv"
	u "build-rest-api/utils"
	"fmt"
)

var CreateContact = func(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.Context())
	user := r.Context().Value("user") . (uint)
	contact := &models.Contact{}
	err := json.NewDecoder(r.Body).Decode(contact)
	if err != nil {
		u.Respond(w, u.Message(false, "Error while decoding request Body"))
		return
	}
	contact.UserId = user
	resp := contact.Create()
	u.Respond(w, resp)
}

var GetContacts = func(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["id"])
	fmt.Println(id)
	fmt.Println(err)
	if err != nil {
		u.Respond(w, u.Message(false, "Something went wrong"))
		return
	}
	data := models.GetContacts(uint(id))
	resp := u.Message(true, "success")
	resp["data"] = data
	u.Respond(w, resp)
}
