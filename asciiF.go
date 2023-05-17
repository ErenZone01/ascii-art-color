package main

import (
	"fmt"
	"os"
	"strings"
)

func foundLine(texte string) []string {
	var compteur int = 0
	var phrase string
	var tabPhrase []string

	for i := 0; i < len(texte); i++ {
		if texte[i] != '\n' {
			phrase += string(texte[i])
		} else {
			compteur++
			if compteur == 9 {
				tabPhrase = append(tabPhrase, phrase)
				phrase = ""
				compteur = 0
				continue
			}
			if compteur != 9 {
				phrase += "\n"
			}

		}

	}
	return tabPhrase
}

func foundposition(position []int, tabPhrase []string) []string {
	var textePhrase []string
	for i := 0; i < len(position); i++ {
		for j := 0; j < len(tabPhrase); j++ {
			if position[i] == j {
				textePhrase = append(textePhrase, tabPhrase[j])
			}
		}
	}
	return textePhrase
}

func compraison(tabASCII []rune, arg1 string, position []int, tab []int) []int {
	var actif bool = false
	var special bool = true

	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(tabASCII); j++ {
			if i < len(arg1)-1 && (string(arg1[i]) == "\\" && string(arg1[i+1]) == "n") {
				actif = true
				break
			} else {
				if actif == false {
					if rune(arg1[i]) == tabASCII[j] {
						position = append(position, tab[j])
						special = true
						break
					} else {
						special = false
					}
				} else {
					actif = false
					break
				}

			}

		}
		if !special {
			error()
		}
	}
	return position
}

func HorzontalLine(x [9][]string) string {
	if IsColor(os.Args[1]) {
		var texte2 string

		for i := 0; i < len(x); i++ {
			for j := 0; j < len(x[i]); j++ {
				texte2 += x[i][j]
			}
			if i > 0 {
				texte2 += "\n"
			}
		}
		return texte2
	} else {
		var texte2 string
		for i, ligne := range x {
			if i > 0 {
				for _, part := range ligne {
					if len(part) > 0 {
						part = part[:len(part)-1]
					}
					texte2 += part
					texte2 += " "

				}
				if len(ligne[0]) == 1 && ligne[0] != "\n" {
					continue
				} else {
					texte2 += "\n"
				}
			}
		}
		return texte2
	}

}

func foundWordColor(arg1, arg2 string) []int {
	var position []int
	for i := 0; i < len(arg1); i++ {
		for j := 0; j < len(arg2); j++ {
			if arg1[i] == arg2[j] {
				position = append(position, j)
			}
		}
	}
	return position
}

func SplitTexte(arg4, arg1, arg3 string, textePhrase []string) [9][]string {
	x := [9][]string{}
	var position []int

	if IsColor(os.Args[1]) {
		position = foundWordColor(arg4, arg1)

		if len(os.Args) == 3 {
			for i := 0; i < len(arg1); i++ {
				position = append(position, i)
			}
		}
		IsCouleur(arg3)
	}

	for j, text := range textePhrase {
		for i, y := range strings.Split(text, "\n") {
			for i := 0; i < len(position); i++ {
				if j == position[i] {
					y = couleur(y, arg3)
				}
			}
			x[i] = append(x[i], y)
		}
	}
	return x
}

func ajoutASCII() []rune {
	var tabASCII []rune
	for i := ' '; i <= '~'; i++ {
		tabASCII = append(tabASCII, i)
	}
	return tabASCII
}

func ajoutNmbr(arg1 string) []int {
	var tab []int
	for i := 0; i < 95; i++ {
		tab = append(tab, i)
	}
	return tab
}

func CaseResolved(arg1 string, arg3 string) bool {
	var actif bool = true
	var text string
	if len(os.Args) >= 2 && arg1 == "\\n" {
		if len(os.Args) == 4 || (IsFile(os.Args[1]) == true) {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(1)
			}
			file.WriteString("\n")
			os.Exit(0)
		}
		fmt.Println()
		return true
	}
	for i := 0; i < len(arg1); i = i + 2 {
		if (i+1 != len(arg1)) && (arg1[i] == '\\' && arg1[i+1] == 'n') {
			actif = false
		} else {
			actif = true
			break
		}

	}
	if actif == false {
		for i := 0; i < len(arg1); i = i + 2 {
			text += "\n"
		}
		if len(os.Args) == 4 || (IsFile(os.Args[1]) == true) {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(0)
			}
			file.WriteString(text)
			os.Exit(0)
		}
		for i := 0; i < len(arg1); i = i + 2 {
			fmt.Println()
		}
		actif = true
		return true
	}

	if len(os.Args) >= 2 && arg1 == "" {
		if len(os.Args) == 4 || (IsFile(os.Args[1]) == true) {
			var file *os.File = writeFile(arg3)
			if file == nil {
				os.Exit(1)
			}
			file.WriteString("")
		}
		return true
	}

	return false
}

func DefineArg() string {
	var arg2 string
	if len(os.Args) == 2 {
		arg2 = "standard"
	} else if len(os.Args) == 3 && IsFile(os.Args[1]) {
		if os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy" {
			arg2 = "standard"
			return arg2
		} else {
			arg2 = os.Args[2]
		}
		return arg2
	} else if (len(os.Args) == 3 || len(os.Args) == 4) && (IsColor(os.Args[1]) == true) {
		arg2 = "standard"
		return arg2
	} else if len(os.Args) == 4 && (IsFile(os.Args[1])) {
		arg2 = os.Args[3]
		if arg2 != "standard" && arg2 != "shadow" && arg2 != "thinkertoy" {
			errorOutput()
		}
		return arg2
	} else if !IsColor(os.Args[1]) && !IsFile(os.Args[1]) && len(os.Args) == 3 {
		arg2 = os.Args[2]
	} else if !IsColor(os.Args[1]) && !IsFile(os.Args[1]) && (len(os.Args) == 3 || len(os.Args) == 4) && os.Args[len(os.Args)-1] != "standard" && os.Args[len(os.Args)-1] != "shadow" && os.Args[len(os.Args)-1] != "thinkertoy" {
		error()
	}
	return arg2
}

func foundFile(arg2 string) string {
	file, err := os.ReadFile(arg2 + ".txt")
	if err != nil {
		error()
	}

	return string(file)
}
func error() {
	fmt.Println("Usage: go run . [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . something standard")
	os.Exit(0)
}
func errorOutput() {
	fmt.Println("Usage: go run . [OPTION] [STRING] [BANNER]")
	fmt.Println()
	fmt.Println("EX: go run . --output=<fileName.txt> something standard")
	os.Exit(0)
}
func errorColor() {
	fmt.Println("Usage: go run . [OPTION] [STRING] ")
	fmt.Println()
	fmt.Println("EX: go run . --color=<color> <letters to be colored> \"something\" ")
	os.Exit(0)
}

func IsFile(arg3 string) bool {
	var option1 = "--output"
	var option = "--output="
	var fichier string
	if len(arg3) >= 8 {
		if option1 == arg3[0:8] {
			if len(arg3) >= 9 {
				if option == arg3[0:9] {
					if len(os.Args) > 2 {
						for i := 9; i < len(arg3); i++ {
							fichier += string(arg3[i])
						}
						if len(fichier) > 4 && fichier[len(fichier)-1] == 't' && fichier[len(fichier)-2] == 'x' && fichier[len(fichier)-3] == 't' && fichier[len(fichier)-4] == '.' {
							return true
						} else {
							errorOutput()
						}
					} else {
						errorOutput()
					}
				} else {
					errorOutput()
				}
			} else {
				errorOutput()
			}
		} else {
			return false
		}
	} else {
		return false
	}
	return false
}
func IsColor(arg3 string) bool {
	option := "--color"
	option2 := "--color="
	if len(arg3) >= 7 {
		if arg3[0:7] == option {
			if len(arg3) > 8 {
				if option2 == arg3[0:8] {
					if len(os.Args) >= 3 {
						return true
					} else {
						errorColor()
					}
				} else {
					errorColor()
				}

			} else {
				errorColor()
			}
		}
	}

	return false
}

func writeFile(arg3 string) *os.File {
	var option = "--output="
	var fichier string
	var file *os.File
	if len(arg3) > 10 {
		if strings.Contains(option, arg3[0:9]) {
			for i := 9; i < len(arg3); i++ {
				fichier += string(arg3[i])
			}
			if len(fichier) > 4 && fichier[len(fichier)-1] == 't' && fichier[len(fichier)-2] == 'x' && fichier[len(fichier)-3] == 't' && fichier[len(fichier)-4] == '.' {
				file, _ = os.Create(fichier)
			} else {
				error()
			}

		} else {
			error()
		}
	} else {
		error()
	}

	return file
}

func IsCouleur(arg3 string) bool {
	var argument = arg3[8:]
	var option = []string{"green", "yellow", "blue", "red", "magenta", "black", "orange"}
	var actif = false
	for i := 0; i < len(option); i++ {
		if argument == option[i] {
			actif = true
			break
		} else {
			actif = false
		}
	}

	if !actif {
		fmt.Print("color not found")
		os.Exit(0)
	}

	return actif
}

func couleur(text, arg3 string) string {
	var option = []string{"green", "yellow", "blue", "red", "magenta", "black", "orange"}
	var argument = arg3[8:]
	for i := 0; i < len(option); i++ {
		if option[i] == argument {
			if option[i] == "green" {
				green := "\033[32m" + text + "\033[0m"
				return green
			} else if option[i] == "red" {
				red := "\033[31m" + text + "\033[0m"
				return red
			} else if option[i] == "yellow" {
				yellow := "\033[33m" + text + "\033[0m"
				return yellow
			} else if option[i] == "blue" {
				blue := "\033[34m" + text + "\033[0m"
				return blue
			} else if option[i] == "black" {
				black := "\033[30m" + text + "\033[0m"
				return black
			} else if option[i] == "magenta" {
				magenta := "\033[35m" + text + "\033[0m"
				return magenta
			} else if option[i] == "orange" {
				magenta := "\033[38;5;208m" + text + "\033[0m"
				return magenta
			}
		} else {
			if i == len(option)-1 {
				errorColor()
			}
		}
	}

	return text
}

func asciiArt(arg1 string, arg3 string, arg4 string, texte string) {
	var arguments = strings.Split(arg1, "\\n")
	var texte2 string
	for i := 0; i < len(arguments); i++ {
		if len(arguments[i]) == 0 {
			texte2 += "\n"
			continue
		}
		var tab []int = ajoutNmbr(arg1)
		var tabPhrase []string
		var tabASCII []rune = ajoutASCII()
		var position []int
		var textePhrase []string

		x := [9][]string{}

		if arguments[i] != "\n" {
			arg1 = arguments[i]
			tab = ajoutNmbr(arg1)

			position = compraison(tabASCII, arg1, position, tab)
			tabPhrase = foundLine(texte)
			textePhrase = foundposition(position, tabPhrase)
			x = SplitTexte(arg4, arg1, arg3, textePhrase)
			texte2 += HorzontalLine(x)

		}

	}
	if len(os.Args) == 4 || len(os.Args) == 3 {
		if IsFile(arg3) {
			var file *os.File = writeFile(arg3)
			file.WriteString(texte2)
			os.Exit(0)
		} else {
			if len(os.Args) == 4 && IsFile(arg3) {
				var file *os.File = writeFile(arg3)
				file.WriteString(texte2)
				os.Exit(0)
			}

		}

	}
	fmt.Print(texte2)
}

