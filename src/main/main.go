package main

import (
	"server"
	"json"
	"git"
)

func main(){
	messages := make(chan string)
	//go server.StartServer(messages)
	server.StartServer(messages)
	println("Waiting for messages")

	for{
		message := <- messages
		//Do whatever has to be done when receiving a message
		repo := json.Decode(message)
		git.Update(repo)
	}

}
