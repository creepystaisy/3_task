package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	//length = len(lines)
	return lines, scanner.Err()
}
func workWithString(a string) string {
	s := strings.Split(a, "/")
	str := strings.Join(s, "")
	s = strings.Split(str, "https:")
	a = strings.Join(s, "")
	fmt.Print(a)
	return a
}
func writeLines(lines []string, path string) /*error*/ {
	var makedDir string

	for i := 0; i < len(lines); i++ {
		resp, err := http.Get(lines[i])
		if err != nil {
			fmt.Println(err)
			continue
			//return err
		}
		a := lines[i]
		a = workWithString(a)
		makedDir = path + "/" + a
		os.MkdirAll(makedDir, 0644)
		openedDir := path + "/" + a + "/" + a + ".txt"
		file, err := os.OpenFile(openedDir, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666) //os.Create(path)
		if err != nil {
			fmt.Println(err)
			continue
			//return err
		}
		defer file.Close()
		w := bufio.NewWriter(file)
		io.Copy(w, resp.Body)
		defer resp.Body.Close()

	}

	// for _, line := range lines {
	// 	fmt.Fprintln(w, line)
	// }
	//return w.Flush()
}

func main() {
	//использавать библ флаг
	inputFile := os.Args[1]
	outputFile := os.Args[2]
	lines, err := readLines(inputFile)
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}
	writeLines(lines, outputFile)
	//var text string = outputFile + "res.txt"
	//if err := writeLines(lines, outputFile); err != nil {
	//	log.Fatalf("writeLines: %s", err)
	//}
}
