package main

import (
	"flag"
	"fmt"
	"strings"
)

const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
	Orange = "\033[38;5;208m" // ANSI escape code for orange color
)

func fad(text string, colorCode string, lettersToColor string) string {
	var s int
	fmt.Println("Choose a font 😊")
	fmt.Println("Standard = 1")
	fmt.Println("Thinkertoy = 2")
	fmt.Println("Shadow = 3")

	fmt.Scan(&s)
	var m2 []string
	var m1 []string
	for i := 0; i < len(text); i++ {
		var char rune = rune(text[i]) // Convert byte to rune
		var shouldColor bool
		if strings.ContainsRune(lettersToColor, char) {
			shouldColor = true
		} else {
			shouldColor = false
		}

		switch s {
		case 1:
			m2 = standard(byte(char)) // Pass rune as byte
		case 2:
			m2 = thinkertoy(byte(char)) // Pass rune as byte
		case 3:
			m2 = shadow(byte(char)) // Pass rune as byte
		}

		if shouldColor {
			m2 = colorizeSlice(m2, colorCode)
		}

		if i == 0 {
			m1 = m2
		} else {
			for j := 0; j < len(m1); j++ {
				m1[j] += m2[j]
			}
		}
	}
	coloredText := strings.Join(m1, "\n")
	return coloredText
}

func colorize(text string, colorCode string) string {
	return colorCode + text + Reset
}

func colorizeSlice(slice []string, colorCode string) []string {
	for i := range slice {
		slice[i] = colorCode + slice[i] + Reset
	}
	return slice
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
	case "orange":
		return Orange
	default:
		fmt.Println("Color not found, using default color (red)")
		return Red
	}
}

func main() {
	colorFlag := flag.String("color", "red", "Specify the color (red, green, yellow, blue, orange)")
	flag.Parse()

	if len(flag.Args()) < 2 {
		fmt.Println("Usage: go run . --color=<color> <letters to be colored> <text>")
		return
	}

	colorCode := chooseColor(*colorFlag)
	lettersToColor := flag.Arg(0)
	text := strings.Join(flag.Args()[1:], " ")

	m2 := strings.Split(strings.TrimSpace(text), "\n")
	asciiArt := []string{}
	for _, line := range m2 {
		asciiArt = append(asciiArt, fad(line, colorCode, lettersToColor))
	}
	asciiArtText := strings.Join(asciiArt, "\n")
	fmt.Println(asciiArtText)
}
