package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func postHandler(w http.ResponseWriter, req *http.Request) {
	bodyReader := req.Body

	var bodyRequest []byte

	_, err := bodyReader.Read(bodyRequest)
	if err != nil {
		panic(err)
	}

	action := &Lake{}
	err = json.Unmarshal(bodyRequest, action)
	if err != nil {
		panic(err)
	}

	fmt.Printf("saved the lake data")

}
