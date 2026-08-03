package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func main() {
	listenSocket()
}

func listenSocket() {
	ListenSocket, err := net.Listen("tcp", "127.0.0.1:8080")
	if err != nil {
		log.Fatal("Не удалось открыть слушающий сокет", err)
	}

	for {
		newSocket, err := ListenSocket.Accept()
		if err != nil {
			log.Fatal("Пользовател не удалось подключиться!", err)
		}

		go handleUser(newSocket)
	}
}

func handleUser(socket net.Conn) {
	defer socket.Close()
	n, err := io.WriteString(socket, "OK\n")
	if err != nil {
		fmt.Println("Ошибка при попытке сервера записать в сокет")
	}

	if n != 3 {
		fmt.Println("Сервер записал ответ не полностью")
	}
}
