package main

import (
	"fmt"
	"os"

	"github.com/mmkader85/golang-basic-testing/src/api/providers/locations_provider"
)

func main() {
	country, err := locations_provider.GetCountry("IN")
	if err != nil {
		fmt.Printf("%+v", err)
		os.Exit(1)
	}
	fmt.Println(country)
}
