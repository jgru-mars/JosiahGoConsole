package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
The author of this code is Josiah Groux

This program is made in Go and will ask the user to enter the names of countries in Europe.
The program continues until the user can not enter any more countries, then they can exit the program. Right before the
program closes, it will proceed to print out the list of countries that the user did get correct.
*/

func main() {
	// list of countries found by user
	var countriesFound []string

	//list of the countries that the user can find
	var countries []string
	countries = append(countries, "france", "spain", "andorra", "portugal", "ireland", "iceland", "norway",
		"sweden", "finland", "denmark", "austria", "hungary", "switzerland", "germany", "luxembourg", "belgium",
		"netherlands", "italy", "monaco", "vatican city", "san marino", "slovenia", "slovakia", "croatia", "serbia",
		"belarus", "russia", "latvia", "ukraine", "moldova", "poland", "czech republic", "kosovo", "greece", "albania",
		"romania", "bulgaria", "lithuania", "bosnia and herzegovina", "estonia", "north macedonia", "malta",
		"montenegro", "liechtenstein", "wales", "scotland", "england", "northern ireland")

	// beginning text
	fmt.Println("This is a game where you enter in names of countries in Europe.")
	fmt.Println("The rules are simple: Enter coutnry names and keep going until you can not think of any other European countries. Let us see how many you get!")
	fmt.Println("Type the name of a European country:")

	// creating a console input reader
	consoleReader := bufio.NewReader(os.Stdin)

loop:
	for true {

		// reads the text from the console
		userInput, _ := consoleReader.ReadString('\n')
		userInput = strings.Replace(userInput, "\n", "", -1)

		// if statements for the user input
		// first if statement checks if the user already found the country, if so it will restart the loop
		if contains(countriesFound, userInput) {
			fmt.Println("You found that country already! Try again!")
			fmt.Println("Type the name of a European country:")
			continue loop
		} else {
			// this if statement is a special case for if the user enters in united kingdom
			if userInput == "united kingdom" {
				fmt.Println("The United Kingdom is a collection of many countries so we will add them all!")
				countriesFound = append(countriesFound, "wales", "england", "scotland", "northern ireland")
			} else {
				// This is another special if statement case for isle of man
				if userInput == "isle of man" {
					fmt.Println("I guess we will count this because it is cool. I believe that it may be legal to have duels here.")
					countriesFound = append(countriesFound, "isle of man")
				} else {
					// this if statement checks if the entered country is indeed a European country, if so it will add it to the found list
					if contains(countries, strings.ToLower(userInput)) {
						fmt.Printf("%s is a country in Europe!\n", userInput)
						countriesFound = append(countriesFound, userInput)
					} else {
						// this last else is if the user does not enter a known country, they get to try again by entering the loop
						fmt.Println("That is not a country, try again my fellow human!")
						fmt.Println("Type the name of a European country:")
						continue loop
					}
				}
			}
		}

		// Asking the user if they know any more countries, if they want to list what they found, or want to exit
	yesno:
		for true {
			fmt.Println("Yes or No, do you know anymore countries in Europe? Or if you want to see the Countries you have already then type 'List' without quotations")

			// user input for yes or no
			yesnoInput, _ := consoleReader.ReadString('\n')
			yesnoInput = strings.Replace(yesnoInput, "\n", "", -1)

			switch strings.ToLower(yesnoInput) {

			// if the user types list, it will print out the countries already found and it also adds comas between each country
			case "list":
				fmt.Print(countriesFound[0])
				for i := 1; i < len(countriesFound); i++ {
					fmt.Print(", ")
					fmt.Print(countriesFound[i])
				}
				fmt.Println("")
				fmt.Println("Type the name of a European country:")
				break yesno
			// if the user types yes it will break the yesno loop and enter the earlier loop to continue asking country names
			case "yes":
				fmt.Println("Type the name of a European country:")
				break yesno
			// if no is typed then it exits the main loop and basically ends the program so that it will list the countries found
			case "no":
				break loop
			}
		}

	}

	// the following will list all the countries found by the user, it will also add comas between each so that it is easier to read
	fmt.Println("The countries that you found are:")
	fmt.Print(countriesFound[0])
	for i := 1; i < len(countriesFound); i++ {
		fmt.Print(", ")
		fmt.Print(countriesFound[i])
	}

}

// this function is for searching the list with the userinput in order to check if it is in the country list.
// If it is, then it will return true and allow the program to add it to the countryFound list
func contains(searchList []string, searchTerm string) bool {
	for _, a := range searchList {
		if a == searchTerm {
			return true
		}
	}
	return false
}
