package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Global struct {
	Fichier   string
	UserInput int
	Loop      bool
	Data      []byte
	Err       error
}

func main() {
	var g Global
	var tmp bool

	for !tmp {
		tmp = g.GetFichier()
	}
	for !g.Loop {
		fmt.Println("Que faire ?\n1. Récupérer tout le texte du fichier\n2. Ajouter du texte dans ce fichier\n3.Supprimer tout le contenu du fichier\n4. Remplacer le contenu\n5. Changer de fichier\n6. Quitter")
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
			g.GetFichier()
		case 6:
			fmt.Println("\033[H\033[2J")
			fmt.Println("Exit successful")
			g.Loop = true
		}
	}
}

func (g *Global) GetFichier() bool {
	var tmp string
	fmt.Println("\033[H\033[2J")
	fmt.Print("quel fichier voulez vous modifier: ")
	fmt.Scan(&tmp)
	file, err := os.OpenFile(tmp, os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return false
	}
	g.Fichier = tmp
	return true
}

func (g *Global) GetData() {
	fmt.Println("\033[H\033[2J")
	g.Data, g.Err = ioutil.ReadFile(g.Fichier)
	if g.Err != nil {
		fmt.Println(g.Err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("le fichier contient:", string(g.Data))
}

func (g *Global) WriteData() {
	input := bufio.NewReader(os.Stdin)
	var tmp string
	var text string
	fmt.Println("\033[H\033[2J")
	file, err := os.OpenFile(g.Fichier, os.O_WRONLY|os.O_APPEND, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Que rajouter dans le fichier: ")
	text, _ = input.ReadString('\n')

	_, err = file.WriteString(text)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("votre texte a bien été ajouté au fichier:", tmp)
}

func (g *Global) RemoveData() {
	var tmp string
	fmt.Println("\033[H\033[2J")
	file, err := os.OpenFile(g.Fichier, os.O_WRONLY|os.O_TRUNC, 0600)
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
	input := bufio.NewReader(os.Stdin)
	var text string
	fmt.Println("\033[H\033[2J")
	file, err := os.OpenFile(g.Fichier, os.O_WRONLY|os.O_TRUNC, 0600)
	defer file.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Que rajouter dans le fichier: ")
	text, _ = input.ReadString('\n')

	_, err = file.WriteString(text)
	if err != nil {
		panic(err)
	}
	fmt.Println("\033[H\033[2J")
	fmt.Println("votre texte a bien été remplacé")

}

func GetUserInput() int {
	var loop bool
	var ret int
	for !loop {
		ret = strToNumber()
		if ret < 1 || ret > 6 {
			fmt.Print("veuillez choisir un nombre de 1 a 6: ")
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
