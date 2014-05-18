package main

import (
	"fmt"
)

var (
	foregroundColors = map[string]int{
		"red":     31,
		"green":   32,
		"yellow":  33,
		"blue":    34,
		"magenta": 35,
		"cyan":    36,
	}

	backgroundColors = map[string]int{
		"red":     41,
		"green":   42,
		"yellow":  43,
		"blue":    44,
		"magenta": 45,
		"cyan":    46,
	}
)

func foreBlackText(text string) string {
	return coloredText(text, 30)
}

func backBlackText(text string) string {
	return coloredText(text, 40)
}

func foreColoredText(text string, colorName string) string {
	color, exists := foregroundColors[colorName]
	if exists {
		return coloredText(text, color)
	} else {
		return text
	}
}

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
