package main

import (
	"fmt"
	"strconv"
)

type Game struct {
	NbMatchstick int
	Loop         bool
	Player       int
	NbToRemove   int
	NbMax        int
}

func main() {
	var g Game
	g.InitGame()
	for !g.Loop {
		fmt.Println("c'est au tour du joueur", g.Player)
		fmt.Print("il y a ", g.NbMatchstick, " allumettes restantes, combien voulez-vous en retirer (maximum 3) : ")
		g.NbToRemove = g.GetNbToRemove()
		g.Loop = g.CheckEnd()
		if g.Player == 1 {
			g.Player = 2
		} else {
			g.Player = 1
		}
	}
}

func (g *Game) InitGame() {
	fmt.Print("veuillez choisir un nombre d'allumette: ")
	g.NbMatchstick = initNbMatchstick()
	g.Loop = false
	g.Player = 1
}

func (g *Game) CheckEnd() bool {
	if g.NbMatchstick > g.NbToRemove {
		g.NbMatchstick -= g.NbToRemove
		return false
	}
	if g.Player == 1 {
		fmt.Println("le joueur", g.Player, "a pris la dernière allumette,\nle joueur 2 a gagné")
	} else {
		fmt.Println("le joueur", g.Player, "a pris la dernière allumette,\nle joueur 1 a gagné")
	}
	return true
}

func (g *Game) GetNbToRemove() int {
	var loop bool
	var ret int
	for !loop {
		ret = strToNumber()
		if ret > 3 || ret < 1 {
			fmt.Print("veuillez choisir un nombre entre 1 et 3 :")
		} else {
			loop = true
		}
	}
	return ret
}

func initNbMatchstick() int {
	var loop bool
	var ret int
	for !loop {
		ret = strToNumber()
		if ret < 4 {
			fmt.Print("veuillez choisir un nombre supérieur à 3: ")
		} else {
			loop = true
		}
	}
	return ret
}

func strToNumber() int {
	var tmp string
	var ret int

	fmt.Scan(&tmp)
	for !IsNumeric(tmp) {
		fmt.Print("veuiller insérer un nombre: ")
		fmt.Scan(&tmp)
	}
	ret, _ = strconv.Atoi(tmp)
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
