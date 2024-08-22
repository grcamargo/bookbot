package main

import "fmt"

const bookPath = "/home/grcamargo/workspace/bookbot/books/frankenstein.txt"

func main() {

	wordNumber, err := countBookWords(bookPath)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(wordNumber)

}
