package main

import (
	"errors"
	"fmt"
	"io"
	"net/http"
)

func deleteHandler(w http.ResponseWriter, req *http.Request) {
	err := req.ParseForm()
	checkError(err)

	id := req.Form.Get("id")

	if id == "" {
		panic(errors.New("id value not found in the request"))
	}

	Store[id] = Lake{Id: "id", Name: "name", Area: "area"}

	_, ok := Store[id]
	if !ok {
		_, err := io.WriteString(w, "404 Not Found")
		if err != nil {
			panic(err)
		}
	}
	delete(Store, id)
	fmt.Println("Deleted a lake data")
}
