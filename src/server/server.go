package server

import (
	"net/http"
	"bytes"
)

var messages chan string


func StartServer(msg chan string){
	//Set up handle function for processing method
	http.HandleFunc("/process", handleProcessRequest)
	//Start Server
	go http.ListenAndServe(":8080", nil)

	//Set channel for storing received messages
	messages = msg
}


/**
Handle function for every request with path "/process"
Reads body and writes it into the channel provided to the server
 */
func handleProcessRequest(w http.ResponseWriter, r *http.Request){
	//TODO: Insert logging for better tracking of requests

	//Read body from http request
	buf := new(bytes.Buffer)
	buf.ReadFrom(r.Body)
	body := buf.String()

	//Store body into received message buffer
	messages <- body

	//Send client a success message
	w.WriteHeader(200)
}
