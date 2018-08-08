package main

import (
	"server"
)

func main(){
	messages := make(chan string)
	server.StartServer(messages)
	for{
		message:= <- messages
		print(message)
	}
}
