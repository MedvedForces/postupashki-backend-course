package main

import (
	"os"
	"net/http"
	"fmt"
	"io"
	"strconv"
	"time"
	"strings"
)

func getHeaders(resp *http.Response) string{
	headers := ""
	for name, value := range resp.Header{
		headers += name + ": "
		headers += strings.Join(value, "") + "\n"
	}

	return headers
}

func getAnswer(link string, limitTime float64) (string, error){
	client := http.Client{
		Timeout : time.Duration(limitTime * 1000) * time.Millisecond,
	}

	resp, err := client.Get(link)
	if err != nil{
		//fmt.Println("Ошибка запроса Get", err.Error())
		return "", err
	}

	defer resp.Body.Close()

	status := resp.StatusCode
	answer := "HTTP/1.1 " + strconv.Itoa(status) + " " + http.StatusText(status) + "\n"
	answer += getHeaders(resp)
	body, err := io.ReadAll(resp.Body)
	if err != nil{
		return "", err
	}

	answer += string(body)
	return answer, nil
}

func answerServer(channel chan string, link string, time float64){
	answer, err := getAnswer(link, time)
	if err != nil{
		fmt.Println("Ошибка ответа от сервера", err)
		os.Exit(666)
	}

	channel <- answer
}

func requestsAll(links []string, timeOut float64){
	channel := make(chan string)
	for _, link := range links{
		go answerServer(channel, link, timeOut)
	}

	go func(timeOut float64){
		time.Sleep(time.Duration(timeOut * 1000) * time.Millisecond)
		fmt.Println("Ошибка по всем запросам")
		os.Exit(228)
	}(timeOut)

	headsAndBody := <- channel //записывает первый полученный ответ в канал
	fmt.Println(string(headsAndBody)) // выводит его и main горутина зав раб не дожидаясь ответа от других горутин
}

func helper(lenArgs int){
	if lenArgs != 2 {
		fmt.Println("Некорректная команда!")
		return
	}

	fmt.Println("Справка по использованию утилиты")
}

func requestWithTime(lenArgs int, argsOs []string){
	timeOut, err := strconv.ParseFloat(argsOs[2], 64)
	if err != nil{
		fmt.Println("Некорректный timeout", err.Error())
		return
	}

	requestsAll(argsOs[3:], timeOut)
}

func main(){
	argsOs := os.Args
	lenArgs := len(argsOs)
	if lenArgs < 2{
		fmt.Println("Некорректная команда!")
		return
	}

	if argsOs[1] == "-h" || argsOs[1] == "--help"{
		if lenArgs < 4{
			fmt.Println("Некорректная команда!")
			return
		}

		helper(lenArgs)
		return
	}

	if argsOs[1] == "-t" || argsOs[1] == "--timeout"{
		requestWithTime(lenArgs, argsOs)
		return
	}

	requestsAll(argsOs[1:], 15)
}