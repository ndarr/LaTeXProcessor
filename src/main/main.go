package main

import (
	"server"
	"fmt"
)

func main(){
	print("Hello World")
	server.StartServer()
	fmt.Scanf("Server running...")
}
