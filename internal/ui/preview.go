package ui

import "fmt"

func ShowPreview(originalMessage string, improvedMessage string) {
	fmt.Println()
	fmt.Println("Original message:")
	fmt.Println()
	fmt.Println(originalMessage)

	fmt.Println()
	fmt.Println("AI suggestion:")
	fmt.Println()
	fmt.Println(improvedMessage)
	fmt.Println()
}
