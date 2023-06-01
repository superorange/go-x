package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	listen, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		log.Println("Error listening:", err)
		panic(err)
	}

	for {
		accept, err := listen.Accept()
		if err != nil {
			log.Println("Error accepting: ", err)
			panic(err)
		}
		go handleRequest(accept)

	}

}

func handleRequest(accept net.Conn) {
	defer func(accept net.Conn) {
		err := accept.Close()
		if err != nil {
			log.Println("Error close: ", err)
		}
	}(accept)

	buf := make([]byte, 1024*1024)

	for n, err := accept.Read(buf); err == nil && n != -1; n, err = accept.Read(buf) {
		fmt.Printf("Received data: %v", string(buf[:n]))
	}

	//if err != nil {
	//	log.Println("Error read: ", err)
	//	return
	//}
	//fmt.Printf("Received data: %v", string(buf[:n]))
}
