package server

import (
	"net/http"
)

var messages chan string


func StartServer(_messages chan string){
	http.HandleFunc("/process", handler)
	http.ListenAndServe(":8080", nil)
	println("Listening on 8080")
	messages = _messages
}


func handler(w http.ResponseWriter, r *http.Request){
	println("Request")
	w.WriteHeader(200)
	w.Write([]byte("Hello World"))
	readCloser:= r.Body
	var body []byte
	readCloser.Read(body)
	println(string(body))
	/*messages <- string(body)*/
}
