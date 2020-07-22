package main

import (
	"bufio"
	"fmt"
	"log"
	"net/http"
	"os"
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

func writeLines(lines []string, path string) error {
	for i := 0; i < len(lines); i++ {
		resp, err := http.Get(lines[i])
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		defer resp.Body.Close()
		for true {

			bs := make([]byte, 1014)
			n, err := resp.Body.Read(bs)
			lines[i] = string(bs[:n])

			if n == 0 || err != nil {
				break
			}
		}
	}

	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	w := bufio.NewWriter(file)
	for _, line := range lines {
		fmt.Fprintln(w, line)
	}
	return w.Flush()
}

func main() {
	lines, err := readLines("url.txt")
	if err != nil {
		log.Fatalf("readLines: %s", err)
	}
	for i, line := range lines {
		fmt.Println(i, line)
	}

	if err := writeLines(lines, "res.txt"); err != nil {
		log.Fatalf("writeLines: %s", err)
	}
}
