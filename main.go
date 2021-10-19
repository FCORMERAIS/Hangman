package main 

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
)

func main() {
	begin()
}

func print_word(letter_choose []string, word string) {
	count := 0
	for i:= 0 ; i < len(word) ;i++{
		for k := 0; k < len(letter_choose); k++ {
			if string(word[i]) == string(letter_choose[k]) {
				fmt.Print(string(word[i]))
				count++
			}
		}
		if count == 0 {
			fmt.Print("_")
		}
		count = 0
	}
	fmt.Println(" ")
	fmt.Println(" ")
}

func choose_word() string {
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
			word = word + string(int(s[i]-32))
		}
	}
	list = append(list,word)
	rand.Seed(time.Now().UnixNano())
	return list[rand.Int31n(83)]
}

func show_José(attemps int) {
	if attemps == 10 {
	}else {
		s, err := ioutil.ReadFile("Hangman.txt")
		if err != nil {
			fmt.Printf(err.Error())
		}
		hangman := ""
		attemps++
		for i := attemps*71-71; i < 71*attemps-1; i++ {
			hangman = hangman + string(s[i])
		}
		fmt.Println(hangman)
	}
}

func take_letter(word2 string) []string{
	var tab []string
	rand.Seed(time.Now().UnixNano())
	tab = append(tab,string(word2[rand.Intn(len(word2)-1)]))
	return tab 
}
func win(wordChoose string, group_letter []string) bool {
	count := 0 
	for i := 0; i < len(wordChoose); i++ {
		for k := 0; k < len(group_letter); k++ {
			if group_letter[k] == string(wordChoose[i]) {
				count++
			}
		}
	}
	if count == len(wordChoose) {
		return true
	}else {
		return false
	}
}

func testLetter(letter string, letter_choose[]string) bool{
	for i := 0 ;i < len(letter_choose) ; i++{
		if letter == string(letter_choose[i]) {
			return false
		}
	}
	return true
}

func letterChooseTest (letter string, word string) bool {
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			return true
		}
	}
	return false
}

func rejouer() bool{
	answer := ""
	fmt.Println("voulez-vous refaire une partie ? [Y/N] : ")
	fmt.Scan(&answer)
	if answer == "yes" || answer == "y"||answer == "YES" || answer == "Y"||answer == "Yes" {
		return true
	}else {
		return false
	}
}


func clear() {
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func begin() {
	attemps := 10
	word := choose_word()
	letterUser := take_letter(word)
	letter := ""
	for attemps > 0 && win(word,letterUser) == false{
		clear()
		fmt.Print("voici le mot que vous devez deviner : ")
		print_word(letterUser,word)
		fmt.Print("vous avez ")
		fmt.Print(attemps)
		fmt.Println(" tentatives avant un d'échoué \n \n")
		fmt.Print("les lettres que vous avez utilisez sont : ")
		fmt.Println(letterUser)
		fmt.Println("")
		fmt.Printf(" entrez un caractère :  ")
		fmt.Scan(&letter)
		if 'a' <= rune(letter[0])&& rune(letter[0]) <= 'z' {
			letter = string(int(letter[0]-32))
		}
		if testLetter(letter,letterUser) == false {
			fmt.Println("vous avez déja rentrez cette lettre au par avant \n \n")
			continue
		}else {
			letterUser = append(letterUser,letter)
		}
		if letterChooseTest(letter, word) == false {
			fmt.Println("la lettre que vous avez choisie n'est pas dans le mot \n \n")
			attemps--
			show_José(attemps)
			continue
		}else {
			fmt.Println("vous avez trouvé une lettre de plus ! \n \n")
		}
	}
	if attemps > 0 {
		fmt.Println(" Bravo vous avez trouvé le mot !!! ")
		fmt.Println()
		fmt.Println()
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}else {
		fmt.Println("Mince, José est mort vous n'avez pas su retrouver le mot :'(")
		fmt.Println()
		fmt.Println()
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}
	if rejouer() == true {
		clear()
		begin()
	}else {
		fmt.Println("a bientôt ! :) ")
	}
}
