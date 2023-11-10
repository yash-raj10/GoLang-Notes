package main

import "fmt"

func main() {

	colors := map[string]string{
		"red":   "#ff0000",
		"green": "#4bf745",
		"white": "#ffffff",
	}

	printMap(colors)

	// making a new map without declaring it
	toys := make(map[string]int)
	toys["helicopter"] = 6
	delete(colors, "helicopter") // to del a key pair

	fmt.Println(colors)
}

func printMap(c map[string]string) {
	for color, hex := range c {
		fmt.Println("Hex code for", color, "is", hex)
	}
}
