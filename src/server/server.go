package server

import (
	"net/http"
	"fmt"
)

func StartServer(){
	http.HandleFunc("/process", handler)
	http.ListenAndServe(":8080", nil)
}


func handler(w http.ResponseWriter, r *http.Request){
	readCloser:= r.Body
	var body []byte
	readCloser.Read(body)
	out := string(body)
	fmt.Fprint(w, out)
}
