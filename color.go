package main

import (
	"fmt"
)

var (
	foregroundColors = map[string]int{
		"red":     31,
		"green":   32,
		"yello":   33,
		"blue":    34,
		"magenta": 35,
		"cian":    36,
	}

	backgroundColors = map[string]int{
		"red":     41,
		"green":   42,
		"yello":   43,
		"blue":    44,
		"magenta": 45,
		"cian":    46,
	}
)

func backColoredText(text string, colorName string) string {
	color, exists := backgroundColors[colorName]
	if exists {
		return coloredText(text, color)
	} else {
		return text
	}
}

func coloredScreenName(screenName string) string {
	firstChar := screenName[0]
	return randomColoredText(screenName, int(firstChar))
}

func randomColoredText(text string, seed int) string {
	colorNumber := len(foregroundColors)
	index := seed % colorNumber

	currentIndex := 0
	for _, color := range foregroundColors {
		if currentIndex == index {
			return coloredText(text, color)
		}
		currentIndex++
	}
	return text
}

func coloredText(text string, color int) string {
	return fmt.Sprintf("\033[%dm%s\033[0m", color, text)
}
