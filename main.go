package main

import (
	"flag"
	"fmt"
	"os"
	"strings"
)

const usageMessage = "Usage: go run . --color=<color> <letters to be colored> [STRING]"

func main() {
	colorFlag := flag.String("color", "", "Specify the color using RGB values (e.g., 255,0,0 for red)")
	flag.Parse()

	if *colorFlag == "" {
		fmt.Println(usageMessage)
		os.Exit(1)
	}

	colors := strings.Split(*colorFlag, ",")
	if len(colors) != 3 {
		fmt.Println(usageMessage)
		os.Exit(1)
	}

	red, green, blue := colors[0], colors[1], colors[2]

	stringArg := flag.Arg(0)
	if stringArg == "" {
		fmt.Println(usageMessage)
		os.Exit(1)
	}

	coloredString := colorizeAsciiArt(stringArg, red, green, blue)
	fmt.Println(coloredString)
}

func colorizeAsciiArt(input string, red, green, blue string) string {
	colorCode := fmt.Sprintf("\x1b[38;2;%s;%s;%sm", red, green, blue)
	resetCode := "\x1b[0m"

	var result string

	for _, char := range input {
		asciiArt := getArt(byte(char))
		coloredAsciiArt := colorizeString(strings.Join(asciiArt, "\n"), colorCode, resetCode)
		result += coloredAsciiArt
	}

	return result
}

func colorizeString(input, colorCode, resetCode string) string {
	return fmt.Sprintf("%s%s%s", colorCode, input, resetCode)
}
