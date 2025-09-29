package main

import "fmt"

func main() {
	cards := []string{"Ace of Diamonds","Five of Diamonds"}
	
	for i:=0; i<len(cards);i++{
		fmt.Println(cards[i])
	}
}
