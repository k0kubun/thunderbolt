package main

import (
	"fmt"
)

var (
	colors = []string{
		"red",
		"green",
		"yellow",
		"blue",
		"magenta",
		"cyan",
		"Red",
		"Green",
		"Yellow",
		"Blue",
		"Magenta",
		"Cyan",
	}
	colorsCount = len(colors)

	foregroundColors = map[string]string{
		"red":     "\033[31m",
		"green":   "\033[32m",
		"yellow":  "\033[33m",
		"blue":    "\033[34m",
		"magenta": "\033[35m",
		"cyan":    "\033[36m",
		"Red":     "\033[31m\033[1m",
		"Green":   "\033[32m\033[1m",
		"Yellow":  "\033[33m\033[1m",
		"Blue":    "\033[34m\033[1m",
		"Magenta": "\033[35m\033[1m",
		"Cyan":    "\033[36m\033[1m",
	}

	backgroundColors = map[string]string{
		"red":     "\033[41m",
		"green":   "\033[42m",
		"yellow":  "\033[43m",
		"blue":    "\033[44m",
		"magenta": "\033[45m",
		"cyan":    "\033[46m",
		"Red":     "\033[41m\033[1m",
		"Green":   "\033[42m\033[1m",
		"Yellow":  "\033[43m\033[1m",
		"Blue":    "\033[44m\033[1m",
		"Magenta": "\033[45m\033[1m",
		"Cyan":    "\033[46m\033[1m",
	}
)

func foreBlackText(text string) string {
	return coloredText(text, "\033[30m")
}

func backBlackText(text string) string {
	return coloredText(text, "\033[40m")
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
	seed := 0
	for index, char := range screenName {
		seed += int(char) + index
	}
	return randomColoredText(screenName, seed)
}

func randomColoredText(text string, seed int) string {
	index := seed % colorsCount
	return coloredText(text, foregroundColors[colors[index]])
}

func coloredText(text string, color string) string {
	return fmt.Sprintf("%s%s\033[0m", color, text)
}
