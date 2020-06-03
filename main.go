package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	var wordsList []string

	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			break
		}
		wordsList = append(wordsList, line)
	}

	word := choseRandomWord(wordsList)
	word = word[:len(word)-1]
	var s, soFar string
	fmt.Println(word)
	currentLetter := 0
	var attempts int

	for {
		if soFar == word {
			fmt.Println("You won!")
			break
		}

		fmt.Print(soFar)
		underscores(word, currentLetter)
		_, err := fmt.Scan(&s)
		if err != nil {
			fmt.Println("k bye")
			return
		}

		if isRight(s, word, currentLetter) {
			fmt.Println("right one!")
			soFar += s
			currentLetter++
			attempts++
		} else {
			fmt.Println("wrong letter")
			attempts++
		}

	}
	fmt.Print("Word ", "'", word, "'")
	fmt.Println(" guessed in", attempts, "attempts")
}

func choseRandomWord(list []string) string {
	var res string
	max := len(list) - 1
	res = list[rand.Intn(max)]
	return res
}

func isRight(c, word string, i int) bool {
	if string(word[i]) == c {
		return true
	}
	return false
}

func underscores(word string, curr int) {
	l := len(word) - 1 - curr
	for i := 0; i < l; i++ {
		fmt.Print("_")
	}
}
