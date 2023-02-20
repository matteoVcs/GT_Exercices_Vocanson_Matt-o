package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Global struct {
	UserInput int
	Loop      bool
	Data      []byte
	Err       error
}

func main() {
	var g Global
	for !g.Loop {
		fmt.Println("Que faire ?\n1. Récupérer tout le texte contenu dans un .txt\n2. Ajouter du texte dans ce fichier\n3.Supprimer tout le contenu du fichier\n4. Remplacer le contenu\n5. Quitter")
		g.UserInput = GetUserInput()
		switch g.UserInput {
		case 1:
			g.GetData()
		case 2:
			g.WriteData()
		case 3:
			g.RemoveData()
		case 4:
			g.ReplaceData()
		case 5:
			fmt.Println("\033[H\033[2J")
			fmt.Println("Exit successful")
			g.Loop = true
		}
	}
}

func (g *Global) GetData() {
	var tmp string
	fmt.Println("\033[H\033[2J")
	fmt.Print("de quel fichier voulez vous récupérer le contenu (\"cancel\" pour annuler): ")
	fmt.Scan(&tmp)
	if tmp == "cancel" {
		return
	}
	g.Data, g.Err = ioutil.ReadFile(tmp)
	if g.Err != nil {
		fmt.Println(g.Err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println(string(g.Data))
}

func (g *Global) WriteData() {
	var tmp string
	var text string
	fmt.Println("\033[H\033[2J")
	fmt.Print("dans quel fichier ecrire (\"cancel\" pour annuler): ")
	fmt.Scan(&tmp)
	if tmp == "cancel" {
		return
	}
	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Que rajouter dans le fichier: ")
	fmt.Scanln(&text)

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("\n", text)
	fmt.Println("votre texte a bien été ajouté au fichier:", tmp)
}

func (g *Global) RemoveData() {
	var tmp string
	fmt.Println("\033[H\033[2J")
	fmt.Print("de quel fichier voulez-vous supprimer le contenu ? (\"cancel\" pour annuler): ")
	fmt.Scan(&tmp)
	if tmp == "cancel" {
		return
	}
	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}

	_, err = file.WriteString("")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("Le contenu du fichier:", tmp, "a bien été supprimé")
}

func (g *Global) ReplaceData() {
	var tmp string
	var text string
	var aze int
	fmt.Println("\033[H\033[2J")
	fmt.Print("dans quel fichier ecrire (\"cancel\" pour annuler): ")
	fmt.Scan(&tmp)
	if tmp == "cancel" {
		return
	}
	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Que rajouter dans le fichier: ")
	aze, _ = fmt.Scan(&text)

	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("\n", text)
	fmt.Println(aze)
	fmt.Println("votre texte a bien été remplacé")

}

func GetUserInput() int {
	var loop bool
	var ret int
	for !loop {
		ret = strToNumber()
		if ret < 1 || ret > 5 {
			fmt.Print("veuillez choisir un nombre de 1 a 5: ")
		} else {
			loop = true
		}
	}
	return ret
}

func strToNumber() int {
	var tmp2 = ""
	var ret int

	fmt.Scan(&tmp2)
	for !IsNumeric(tmp2) {
		fmt.Print("veuiller insérer un nombre: ")
		fmt.Scan(&tmp2)
	}
	ret, _ = strconv.Atoi(tmp2)
	return ret
}

func IsNumeric(s string) bool {
	var str = []byte(s)
	for i := 0; i != len(s); {
		if str[i] >= 48 && str[i] <= 57 || str[i] == '-' {
			i++
		} else {
			return false
		}
	}
	return true
}
