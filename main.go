package main 

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	show_José(1)
}
/*
func print_word(letter_choose []rune, word string) {
	mot := ""
	for i:= 0 ; i < len(word) ;i++{
		mot = mot + "_"
	}
	for i:= 0; i < len(letter_choose); i++ {
		for k:= 0 ; k<len(word); k++ {
			if letter_choose[i] == rune(word[k]) {
				mot[k] = letter_choose[i]
			}
		}
	}
	fmt.Println(mot)
}
*/
func choose_word() {
	s, err := ioutil.ReadFile("word1.txt")
	if err != nil {
		fmt.Printf(err.Error())
	}
	var list []string
	var word string = ""
	for i := 0; i < len(s); i++ {
		if string(s[i]) == "\n" {
			list = append(list,word)
			word = ""
		}else {
			word = word + string(s[i])
		}
	}
	list = append(list,word)
	rand.Seed(time.Now().UnixNano())
	fmt.Println(list[rand.Int31n(84)])
}

func show_José(attemps int) {
	s, err := ioutil.ReadFile("Hangman.txt")
	if err != nil {
		fmt.Printf(err.Error())
	}
	hangman := ""
	for i := attemps*71-71; i < 71*attemps-1; i++ {
		hangman = hangman + string(s[i])
	}
	fmt.Println(hangman)
}
