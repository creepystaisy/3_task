package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// this is a comment
func readFile() string { //функция считывающая урлы
	fileRead, err := os.Open("url.txt") //открываем файл с урлами

	if err != nil { // если возникла ошибка
		fmt.Println(err)
		os.Exit(1) // выходим из программы
	}

	defer fileRead.Close()

	data := make([]byte, 64)
	var x string
	for {
		n, err := fileRead.Read(data)
		if err == io.EOF { // если конец файла
			break // выходим из цикла
		}
		//fmt.Print(string(data[:n]))
		//fmt.Println(fileRead.Name()) // просто проверка
		x = string(data[:n])
	} // закрываем файл
	return x
}

func create(name string) {
	file, err := os.Create(name) // создаем файл
	if err != nil {              // если возникла ошибка
		fmt.Println("Unable to create file:", err)
		os.Exit(1) // выходим из программы
	}
	defer file.Close() // закрываем файл
	file.WriteString(parsing(readFile()))

	//fmt.Println(file.Name())               // hello.txt

	//file.WriteString(text)
}

func parsing(url string) string {
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

	}
	defer resp.Body.Close()
	var z string
	for true {

		bs := make([]byte, 1014)
		n, err := resp.Body.Read(bs)
		z = (string(bs[:n]))

		if n == 0 || err != nil {
			break
		}
	}
	return z
}

func main() {
	fmt.Println("hell")
	create("res.txt")

}
