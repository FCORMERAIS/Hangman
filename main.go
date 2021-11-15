package main 

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"time"
	"os"
	"strings"
)

func main() {
	const hangman = `
 _                                             
| |                                            
| |__   __ _ _ __   __ _ _ __ ___   __ _ _ __  
| '_ \ / _  |  _ \ / _  |  _   _ \ / _  |  _ \ 
| | | | (_| | | | | (_| | | | | | | (_| | | | |
|_| |_|\__,_|_| |_|\__, |_| |_| |_|\__,_|_| |_|
                    __/ |                      
                   |___/ 
by Flavio & Dylan
`
	fmt.Print(hangman)
	begin()
}

func printWord(letter_choose []string, word []string) {
	/*
	fonction permettant d'afficher le mot en fonction des lettres que l'utilisateur a déjà trouver
	input : -letter_choose type []string il s'agit des lettres que l'utilisateur a déjà rentrer 
			-word : type string il s'agit du mot a deviner 
	copléxité O(n²)
	*/
	count := 0 
	for i:= 0 ; i < len(word) ;i++{
		for k := 0; k < len(letter_choose); k++ { // debut de la boucle
			if string(word[i]) == string(letter_choose[k]) {
				fmt.Print(string(word[i]))
				count++
			}
		} //fin de la boucle
		if count == 0 {
			fmt.Print("_") // renvoie un underscore si le joueur n'a pas trouver la lettre 
		}
		count = 0
	}
	fmt.Print("\n \n") // saut a la ligne
}

func chooseWord() string {
	/*
	fonction permettant de prendre un mot aléatoire dans une banque de mots 
	return : un mot aléatoire de type sring
	complexité : O(n) ; n = nombre de caractère de tout les mots reunis 
	*/
	s, err := ioutil.ReadFile("word1.txt") // ouverture du fichier word1 contenant tout les mots et qui seront stockés dans s
	if err != nil {
		fmt.Printf(err.Error()) // renvois de l'erreur lors de louverture du fichier si il y a un problème 
		fmt.Println(" Le fichier word1.txt a planter...")
		os.Exit(1)
	}
	list := strings.Split(string(s),"\n") 
	rand.Seed(time.Now().UnixNano()) // ceci permet de faire en sorte que l'aléatoire marche 
	return strings.ToUpper(list[rand.Int31n(83)]) // renvois un mot choisis aléatoirement dans notre liste de mots 
}

func showJosé(attemps int) {
	/*
	fonction permettant d'afficher le pendu 
	input : - attemps type int il sagit du nombre de tentative qu'il reste avant de perde il permet donc d'afficher la position du bon pendu 
	compléxité : O(71)
	*/
	if attemps == 10 { // ne rien renvoyer si il n'y a pas eu d'erreur
	}else {
		s, err := ioutil.ReadFile("Hangman.txt") // ouverture du fichier txt Hangman contenant tout les position possible du pendu et les stock dans "s"
		if err != nil {
			fmt.Printf(err.Error()) // renvois d'une erreur si il y en a une 
			fmt.Println(" Le fichier Hangman.txt a planter...")
			os.Exit(1)
		}
		hangman := ""
		attemps++ // incrémentation de attemps
		for i := attemps*71-71; i < 71*attemps-1; i++ { // debut boucle
			hangman = hangman + string(s[i]) // on ajoute caractèere par caractère le hangman stockés dans s dans la variable string hangman 
		} // fin boucle
		fmt.Println(hangman) // on imprime hangman 
	}
}

func take_letter(word2 []string) []string{
	/*
	fonction permettant de donner une lettre qui est présente dans le mot a deviner 
	input : word2 type string il s'agit du mot a deviner 
	return : List/string il s'agit de la liste des lettres choisie par l'utlisateur 
	compléxité : O(n) ; n = longueur du mot
	*/
	var tab []string
	rand.Seed(time.Now().UnixNano())
	tab = append(tab,string(word2[rand.Intn(len(word2)-1)])) // prend une lettre aléatoire dans le mot a deviner et l'ajoute dans les lettres déja choisie 
	for i := 0; i < len(word2); i++ {
		if string(word2[i]) == "-" { // si il y a un tiret l'ajoute aussi dans notre liste 
			tab = append(tab,"-") 
			return tab 
		}
	}
	return tab // renvois le tableau contenant 
}
func win(wordChoose []string, group_letter []string) bool {
	/*
	fonction qui permet de savoir si l'utilisateur a gagné en trouver toutes les lettres 
	input : -wordChoose type string il sagit du mot que l'utilisateur doit choisir 
			-group_letter type List/string il s'agit des lettres que l'utilisateur a choisi 
	return : bool 
	Complexité : O(n²)
	*/
	count := 0 
	for i := 0; i < len(wordChoose); i++ { // debut boucle i
		for k := 0; k < len(group_letter); k++ { // debut boucle k 
			if group_letter[k] == string(wordChoose[i]) { // on vérifie si une lettre du mot est dans notre liste de mot 
				count++ // si la lettre du mot est dans notre liste de lettre on ajoute 1 a un compteur 
			}// fin obucle k
		} // fin boucle i
	}
	if count == len(wordChoose) { // si notre compteur est égal a la longueur du mot a deviner alors nous avons trouver toutes les lettres
		return true
	}else { // sinon non
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
		if letter == string(letter_choose[i]) { // on vérifie si la lettre choisie par l'utilisateur n'est pas déjà présente dans notre liste de mot si c'estle cas on renvois false
			return false
		}
	}
	return true
}

func letterChooseTest(letter string, word []string) bool {
	/*
	fonction permettant de vérifier si la lettre chosi par l'utilisateur est contenu dans le mot a deviner
	input : -letter : type string il sagit de la lettre choisi par l'utilisateur
			-word : type string il sagit du mot a deviner
	return : Bool
	complexité : O(2n) ; n = len(word)
	*/
	for i := 0; i < len(word); i++ {
		if letter == string(word[i]) { // on vérifie si la lettre choisi par l'utilisateur est présente dans le mot a deviner si c'est le cas en renvois true 
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
	if answer == "yes" || answer == "y"||answer == "YES" || answer == "Y"||answer == "Yes" { // si l'utilisateur ecrit yes il rejoue une partie sinon il quitte le programme 
		return true
	}else {
		return false
	}
}


func clear() {
	/*
	fonction permettant de clear la console pour que l'affichage soit plus propre 
	compléxité : O(30)
	*/
	for i := 0; i < 30; i++ {
		fmt.Println() // retour a la ligne 
	}
}

func begin() {
	/*
	fonction principale du programme il permet de mettre en relation toutes les variables ci dessus, cette fonciton permet de jouer au pendu 
	*/
	attemps := 10 // il s'agit du nombre de tentatives qu'il nous reste 
	word := chooseWord() // on choisi un mot aléatoire
	test := strings.Split(word, "")
	letterUser := take_letter(test)
	letter := ""
	for attemps > 0 && win(test,letterUser) == false{ // on continu de jouer tant que la fonction win est fausse et que le nombre de tentatives est strictement supérieur a 0 /
		clear()
		showJosé(attemps)// imprime la position du pendu
		fmt.Print("voici le mot que vous devez deviner : ")
		printWord(letterUser,test)// imprime le mot que l'on doit deviner avec seulement les lettres 
		fmt.Print("vous avez ")
		fmt.Print(attemps) // on imprime 
		fmt.Print(" tentatives avant un d'échoué \n \n \n")
		fmt.Print("les lettres que vous avez utilisez sont : ")
		fmt.Println(letterUser) // imprime toutes les lettres que l'utilisateur a rentré 
		fmt.Println("")
		fmt.Printf(" entrez un caractère :  ")
		fmt.Scan(&letter)// on prend la lettre uqe l'utilisateur choisie
		fmt.Println()
		letter = strings.ToUpper(letter) // si la lettre est bien dans dans l'alphabet on le passe en majuscule 
		if testLetter(letter,letterUser) == false { // on verifie si la lettre n'a jamais été choisie
			fmt.Print("vous avez déja rentrez cette lettre au par avant \n \n \n") 
			continue
		}else {
			letterUser = append(letterUser,letter) // on ajoute la lettre choisi par l'utilisateur dans letterUser si il ne la pas encore choisie 
		}
		if letterChooseTest(letter, test) == false { // on verifie si la lettre choisie par l'utilisateur est dans le mot ou pas 
			fmt.Print("la lettre que vous avez choisie n'est pas dans le mot \n \n \n")
			attemps-- // si non on imprime que elle n'est pas dans le mot et on décremente de 1 attemps
			continue
		}else { // sinon cela veut dire que la lettre est forcément dans le mot 
			fmt.Print("vous avez trouvé une lettre de plus ! \n \n \n")
		}
	}//FIN BOUCLE
	if attemps > 0 { // si il reste des tentatives cela veut dire que l'utilisateur a gagner 
		fmt.Print(" Bravo vous avez trouvé le mot !!! \n \n \n")
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}else { // sinon il a perdu
		fmt.Print("Mince, José est mort vous n'avez pas su retrouver le mot :'( \n \n \n")
		fmt.Print("le mot était : ")
		fmt.Println(word)
	}
	if replay() == true { // si replay est égal a true on relance le programme begin qui recommence une partie 
		clear()
		begin()
	}else { //sinon le programme s'arrete
		fmt.Println("a bientôt ! :) ")
	}
}