package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"strconv"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", "localhost:9090")
	checkError(err)
	for {
		conn, err := listener.Accept() // принимаем TCP-соединение от клиента и создаем новый сокет
		checkError(err)
		go handleClient(conn) // обрабатываем запросы клиента в отдельной горутине
	}
}

func handleClient(conn net.Conn) {
	var inputNumber int
	var endOfByte int
	rand.Seed(time.Now().UnixNano())
	number := rand.Intn(100)
	fmt.Println("\nЗагаданно новое число - ", number)
	buf := make([]byte, 16) // буфер для чтения клиентских данных
	for {
		_, err := conn.Read(buf) // читаем из сокета
		for i, v := range buf {
			if v == '\n' {
				endOfByte = i
				break
			}
		}
		if string(buf[:5]) == "QUESS" {
			inputNumber, err = strconv.Atoi(string(buf[6:endOfByte]))
			checkError(err)
			fmt.Println("Получено число: ", inputNumber)
			if inputNumber > number {
				conn.Write([]byte("MORE\n"))
				checkError(err)
			} else if inputNumber < number {
				conn.Write([]byte("LESS\n"))
				checkError(err)
			} else {
				conn.Write([]byte("EQUAL\n"))
				checkError(err)
				break
			}
		} else if string(buf[:4]) == "EXIT" {
			conn.Write([]byte("Goodbye\n"))
			checkError(err)
		}
	}
}

func checkError(err error) {

	if err != nil {
		log.Fatal(err)
	}
}
