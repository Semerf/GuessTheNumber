package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"time"
)

func main() {
	conn, _ := net.Dial("tcp", ":9090")
	var quess_str string = "QUESS %d\n"
	var query_str string
	var max int = 100
	var min int = 0
	rand.Seed(time.Now().UnixNano())
	var num int
	for {
		num = min + rand.Intn(max-min+1)
		query_str = fmt.Sprintf(quess_str, num)
		// Отправляем в socket
		fmt.Fprint(conn, query_str)
		// Прослушиваем ответ
		message, _ := bufio.NewReader(conn).ReadString('\n')
		if message == "MORE\n" {
			fmt.Println("Ваше число -", num, ".Загаданное число меньше")
			max = num - 1
		} else if message == "LESS\n" {
			fmt.Println("Ваше число -", num, ".Загаданное число больше")
			min = num + 1

		} else if message == "EQUAL\n" {
			fmt.Println("Вы угадали загаданное число -", num, "!")
			break
		} else {
			fmt.Println("Найдено неожиданное значение", message)
		}

	}
}
