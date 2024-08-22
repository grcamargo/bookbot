package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func countBookWords(pathToBook string) (int, error) {
	f, err := os.Open(pathToBook)
	if err != nil {
		log.Fatalf("unable to read file: %v", err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	numberWords := 0

	for scanner.Scan() {
		numberWords++
	}

	if err := scanner.Err(); err != nil {
		return 0, err
	}

	buf := make([]byte, 2)
	for {
		n, err := f.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println(err)
			continue
		}
		if n > 0 {
			fmt.Println(string(buf[:n]))

		}
	}
	return numberWords, nil
}
