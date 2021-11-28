package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
	}
	fmt.Println(colors)
	new_colors := make(map[int]string)
	new_colors[10] = "#ffffff"
	delete(new_colors, 10)
	fmt.Println(new_colors)
}
