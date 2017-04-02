package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	adjectives := readWordlistFile("adjectives.txt")
	nouns := readWordlistFile("nouns.txt")
	parts := []string{randomWord(adjectives), randomWord(nouns), randomNumber()}

	fmt.Println(strings.Join(parts, "-"))
}

func readWordlistFile(filepath string) []string {
	var wordlist []string
	wordlistFile, err := os.Open(filepath)
	if err != nil {
		panic(err)
	}

	defer wordlistFile.Close()

	scanner := bufio.NewScanner(wordlistFile)

	for scanner.Scan() {
		wordlist = append(wordlist, scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return wordlist
}

func randomWord(wordlist []string) string {
	rand.Seed(time.Now().UnixNano())

	return wordlist[rand.Intn(len(wordlist))]
}

func randomNumber() string {
	return strconv.Itoa(rand.Intn(10000))
}
