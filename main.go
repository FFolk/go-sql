package main

import (
	"fmt"
	"go-sql/database"
)

func main() {
	// country, err := database.GetCountryByID(1)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", country)

	// pakchong := database.Country{}
	// pakchong.SetName("ปากช่อง")
	// aff, err := database.AddCountry(&pakchong)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("%v\n", aff)

	aff, err := database.DeleteCountry(133)
	if err != nil {
		panic(err)
	}
	fmt.Println(aff)

	// country := database.Country{}
	// country.SetIdx(111)
	// country.SetName("ประเทศปากช่อง")

	// update, err := database.UpdateCountry(&country)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println(update)

	countries, err := database.GetCountry()
	if err != nil {
		panic(err)
	}
	fmt.Printf("%v\n", countries)
}
