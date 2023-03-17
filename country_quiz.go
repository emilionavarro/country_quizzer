package main

import (
	"fmt"
	"github.com/fatih/color"
)

var countriesAndCapitals = map[string]string{
	"Albania":         "Tirana",
	"Andorra":         "Andorra la Vella",
	"Austria":         "Vienna",
	"Belarus":         "Minsk",
	"Belgium":         "Brussels",
	"Bosnia and Herzegovina": "Sarajevo",
	"Bulgaria":        "Sofia",
	"Croatia":         "Zagreb",
	"Cyprus":          "Nicosia",
	"Czech Republic":  "Prague",
	"Denmark":         "Copenhagen",
	"Estonia":         "Tallinn",
	"Finland":         "Helsinki",
	"France":          "Paris",
	"Germany":         "Berlin",
	"Greece":          "Athens",
	"Hungary":         "Budapest",
	"Iceland":         "Reykjavik",
	"Ireland":         "Dublin",
	"Italy":           "Rome",
	"Kosovo":          "Pristina",
	"Latvia":          "Riga",
	"Liechtenstein":   "Vaduz",
	"Lithuania":       "Vilnius",
	"Luxembourg":      "Luxembourg City",
	"Malta":           "Valletta",
	"Moldova":         "Chisinau",
	"Monaco":          "Monaco",
	"Montenegro":      "Podgorica",
	"Netherlands":     "Amsterdam",
	"North Macedonia": "Skopje",
	"Norway":          "Oslo",
	"Poland":          "Warsaw",
	"Portugal":        "Lisbon",
	"Romania":         "Bucharest",
	"Russia":          "Moscow",
	"San Marino":      "San Marino",
	"Serbia":          "Belgrade",
	"Slovakia":        "Bratislava",
	"Slovenia":        "Ljubljana",
	"Spain":           "Madrid",
	"Sweden":          "Stockholm",
	"Switzerland":     "Bern",
	"Ukraine":         "Kyiv",
	"United Kingdom":  "London",
	"Vatican City":    "Vatican City",
}

func main() {
	fmt.Println("Welcome to the Europe Countries and Capitals Quiz!")
	fmt.Println("You will be presented with a country, and you must guess its capital.")
	fmt.Println("Type 'quit' to end the quiz.")
	fmt.Println()

	var correct, total int

	for country, capital := range countriesAndCapitals {
		total++
		fmt.Printf("What is the capital of %s? ", country)

		var answer string
		fmt.Scanln(&answer)

		if answer == "quit" {
			break
		}

		if answer == capital {
			//fmt.Println("Correct!")
			color.Green("Correct!")
			correct++
		} else {
			color.Red("Incorrect! The capital of %s is %s.\n", country, capital)
		}
	}

	fmt.Printf("You got %d out of %d correct!\n", correct, total)
}
