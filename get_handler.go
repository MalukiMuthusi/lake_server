package main

import (
	"encoding/json"
	"errors"
	"io"
	"net/http"
)

// handle request for getting a lake data from the store
func getHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	checkError(err)

	id := req.Form.Get("id")

	if id == "" {
		panic(errors.New("id value not found in the request"))
	}
	Store[id] = Lake{Id: "id", Name: "name", Area: "area"}

	// check the store for the lake with the ID in the payload
	payload, ok := Store[id]
	if !ok {
		_, err := io.WriteString(w, "404 Not Found")
		if err != nil {
			panic(err)
		}
		return
	}

	// respond by returning the lake found
	responseData, err := json.Marshal(payload)
	if err != nil {
		panic(err)
	}

	_, err = w.Write(responseData)
	if err != nil {
		panic(err)
	}
}
