package main

import (
	"fmt"
	"os"
)

func main() {
	for {
		// Display menu options
		fmt.Println("Val:")
		fmt.Println("1. Skapa device")
		fmt.Println("2. Lista alla")
		fmt.Println("3. Ändra device")
		fmt.Println("4. Sök")
		fmt.Println("5. Avsluta")

		var choice int

		fmt.Print("Enter youe choice: ")
		_, err := fmt.Scanln(&choice)

		if err != nil {
			fmt.Println("Invalid choice. Please try again.")
			continue
		}
		switch choice {

		case 1:
			fmt.Println("You selected Option 1")
		case 2:
			fmt.Println("You selected Option 2")
		case 3:
			fmt.Println("You selected Option 3")
		case 4:
			fmt.Println("You selected Option 4")
		case 5:
			fmt.Println("Exiting...")
			os.Exit(0)
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}
