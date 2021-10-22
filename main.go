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

func printWord(letter_choose []string, word string) {
	/*
	fonction permettant d'afficher le mot en fonction des lettres que l'utilisateur a déjà trouver

	input : -letter_choose type []string il s'agit des lettres que l'utilisateur a déjà rentrer 
			-word : type string il s'agit du mot a deviner 

	copléxité O(n²)
	*/
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

func chooseWord() string {
	/*
	fonction permettant de prendre un mot aléatoire dans une banque de mots 
	*/
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

func showJosé(attemps int) {
	/*
	fonction permettant d'afficher le pendu 

	input : - attemps type int il sagit du nombre de tentative qu'il reste avant de perde il permet donc d'afficher la position du bon pendu 

	compléxité : O(71)
	*/
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
	/*
	fonction permettant de donner une lettre qui est présente dans le mot a deviner 

	input : word2 type string il s'agit du mot a deviner 

	return : List/string il s'agit de la liste des lettres choisie par l'utlisateur 

	compléxité : O(n) ; n = longueur du mot
	*/
	var tab []string
	rand.Seed(time.Now().UnixNano())
	tab = append(tab,string(word2[rand.Intn(len(word2)-1)]))
	for i := 0; i < len(word2); i++ {
		if string(word2[i]) == "-" {
			tab = append(tab,"-") 
			return tab 
		}
	}
	return tab 
}
func win(wordChoose string, group_letter []string) bool {
	/*
	fonction qui permet de savoir si l'utilisateur a gagné en trouver toutes les lettres 

	input : -wordChoose type string il sagit du mot que l'utilisateur doit choisir 
			-group_letter type List/string il s'agit des lettres que l'utilisateur a choisi 

	return : bool 

	Complexité : O(n²)
	*/
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
	/*
	fonction permettant de vérifier si la lettre choisi par l'utilisateur est déjà contenu dans la liste des lettres qu'à choisi l'utilisateur 

	input : -letter type string il s'agit de la lettre choisi par l'utilsateur
			-letter_choose type List/string il s'agit de la liste de lettres qu'à déjà rentré l'utilisateur 

	return : Bool

	Compléxité : O(2n) ; n = letter_choose
	*/
	for i := 0 ;i < len(letter_choose) ; i++{
		if letter == string(letter_choose[i]) {
			return false
		}
	}
	return true
}

func letterChooseTest(letter string, word string) bool {
	/*
	fonction permettant de vérifier si la lettre chosi par l'utilisateur est contenu dans le mot a deviner

	input : -letter : type string il sagit de la lettre choisi par l'utilisateur
			-word : type string il sagit du mot a deviner

	return : Bool

	complexité : O(2n) ; n = len(word)
	*/
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) {
			return true
		}
	}
	return false
}

func replay() bool{
	/*
	fonction permettant de savoir si l'utilisateur veut relancer une partie ou non 

	return : Bool

	compléxité : O(4)
	*/
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
	/*
	fonction permettant de clear la console pour que l'affichage soit plus propre 

	O(30)
	*/
	for i := 0; i < 30; i++ {
		fmt.Println()
	}
}

func begin() {
	/*
	fonction principale du programme il permet de mettre en relation toutes les variables ci dessus, cette fonciton permet de jouer au pendu 
	*/
	attemps := 10
	word := chooseWord()
	letterUser := take_letter(word)
	letter := ""
	for attemps > 0 && win(word,letterUser) == false{
		clear()
		showJosé(attemps)
		fmt.Print("voici le mot que vous devez deviner : ")
		printWord(letterUser,word)
		fmt.Print("vous avez ")
		fmt.Print(attemps)
		fmt.Print(" tentatives avant un d'échoué \n \n \n")
		fmt.Print("les lettres que vous avez utilisez sont : ")
		fmt.Println(letterUser)
		fmt.Println("")
		fmt.Printf(" entrez un caractère :  ")
		fmt.Scan(&letter)
		fmt.Println()
		if 'a' <= rune(letter[0])&& rune(letter[0]) <= 'z' {
			letter = string(int(letter[0]-32))
		}
		if testLetter(letter,letterUser) == false {
			fmt.Print("vous avez déja rentrez cette lettre au par avant \n \n \n")
			continue
		}else {
			letterUser = append(letterUser,letter)
		}
		if letterChooseTest(letter, word) == false {
			fmt.Print("la lettre que vous avez choisie n'est pas dans le mot \n \n \n")
			attemps--
			continue
		}else {
			fmt.Print("vous avez trouvé une lettre de plus ! \n \n \n")
		}
	}
	if attemps > 0 {
		fmt.Print(" Bravo vous avez trouvé le mot !!! \n \n \n")
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}else {
		fmt.Print("Mince, José est mort vous n'avez pas su retrouver le mot :'( \n \n \n")
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}
	if replay() == true {
		clear()
		begin()
	}else {
		fmt.Println("a bientôt ! :) ")
	}
}
