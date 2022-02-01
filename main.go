package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode/utf8"
)

func main() {
	mapWord := createWordMap()
	searchInWordMap(mapWord)
}

func censor(word string) {
	censored := strings.Repeat("*", utf8.RuneCountInString(word))
	fmt.Println(censored)
}

func searchInWordMap(mapWord map[string]struct{}) {
	var word string
	x := true
	for x == true {
		fmt.Scan(&word)
		if _, ok := mapWord[strings.ToLower(word)]; ok {
			censor(word)
		} else if strings.ToLower(word) == "exit" {
			x = false
			fmt.Println("Bye")
		} else {
			fmt.Println(word)
		}
	}
}

func createWordMap() map[string]struct{} {
	var filename string
	fmt.Scan(&filename)
	mapWord := make(map[string]struct{})
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		mapWord[scanner.Text()] = struct{}{}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return mapWord
}
