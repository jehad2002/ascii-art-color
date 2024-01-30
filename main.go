package main

import (
	"fmt"
	"os"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	// Add more colors as needed
)

func fad(text string, colorCode string) string {
	var s int
	fmt.Println("Choose a font 😊")
	fmt.Println("Standard = 1")
	fmt.Println("Thinkertoy = 2")
	fmt.Println("Shadow = 3")

	fmt.Scan(&s)
	var m2 []string
	var m1 []string
	for i := 0; i < len(text); i++ {
		switch s {
		case 1:
			m2 = standard(text[i])
		case 2:
			m2 = thinkertoy(text[i])
		case 3:
			m2 = shadow(text[i])
		}
		if i == 0 {
			m1 = m2
		} else {
			for j := 0; j < len(m1); j++ {
				m1[j] += m2[j]
			}
		}
	}
	coloredText := colorize(strings.Join(m1, "\n"), colorCode)
	return coloredText
}

func colorize(text string, colorCode string) string {
	return colorCode + text + Reset
}

func chooseColor(colorFlag string) string {
	switch colorFlag {
	case "red":
		return Red
	case "green":
		return Green
	case "yellow":
		return Yellow
	case "blue":
		return Blue
	default:
		fmt.Println("the color not found, using default color")
		return colorFlag
	}
}

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . --color=<color> <text>")
		return
	}

	colorFlag := os.Args[1][8:]
	colorCode := chooseColor(colorFlag)
	text := os.Args[2]

	m2 := strings.Split(strings.TrimSpace(text), "\n")
	asciiArt := []string{}
	for _, line := range m2 {
		asciiArt = append(asciiArt, fad(line, colorCode))
	}
	asciiArtText := strings.Join(asciiArt, "\n")
	fmt.Println(asciiArtText)
}
