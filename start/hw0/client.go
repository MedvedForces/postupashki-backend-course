package main

import (
	"fmt"
	"io"
	"net"
	"os"
)

func main() {
	ans, err := client()
	if err != nil {
		fmt.Println("Ошибка при чтении в клиенте с сервера", err)
		os.Exit(228)
	}

	if ans == "OK\n" {
		fmt.Println("Все хорошо ответ верный")
	} else {
		fmt.Println("Сервер выдал неверный ответ!")
		os.Exit(666)
	}
}

func client() (string, error) {
	servAdress := "127.0.0.1:8080"
	conn, err := net.Dial("tcp", servAdress)
	if err != nil {
		fmt.Println("Ошибка при подключении к серверу на 127.0.0.1:8080")
	}

	return handleServer(conn)
}

func handleServer(conn net.Conn) (string, error) {
	defer conn.Close()

	buffer, err := io.ReadAll(conn)
	return string(buffer), err
}
