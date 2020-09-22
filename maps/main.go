package main

import "fmt"

func main() {
	countryMap := make(map[string]string)

	countryMap["Zimbabwe"] = "Harare"
	countryMap["Kenya"] = "Nairobi"
	countryMap["Uganda"] = "Kampala"
	countryMap["Nigeria"] = "Abuja"
	countryMap["Libya"] = "Tripoli"

	fmt.Printf("There are %d entries in the country map\n", len(countryMap))

	for country, capital := range countryMap {
		fmt.Printf("The capital of %s is %s\n", country, capital)
	}

	delete(countryMap, "Uganda")
	fmt.Println(countryMap)
	fmt.Printf("There are %d entries in the country map\n", len(countryMap))
}
