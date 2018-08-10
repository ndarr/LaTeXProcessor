package main

import "server"

func main(){
	messages := make(chan string)
	//go server.StartServer(messages)
	server.StartServer(messages)
	println("Waiting for messages")

	for{
		message := <- messages
		//Do whatever has to be done when receiving a message
		println(message)
	}

}
