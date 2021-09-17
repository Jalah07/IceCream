package main

// Import the necessary packages and call the appropriate methods to read a files input and convert it to an IceCream
// You should then validate IceCream if it fails validation print that out.
// If the icecream passes validation, store the IceCream in the db
import (
	//"fmt"
	"ReadFileProj/FileReader"
	"ReadFileProj/Validation"

	//"ReadFileProj/ice_cream"
	"fmt"

	"ReadFileProj/Repo"
)

func main() {
	files := []string{"ic1", "ic2", "ic3", "ic4"}
	for _, v := range files {
		fn := "input/" + v + ".json"

		iceCream, err := FileReader.ReadFile(fn)
		if err != nil {
			fmt.Println(err)
			return
		}
		//fmt.Println("******************************")
		//fmt.Printf("%+v\n", iceCream)

		fmt.Printf("%+v\n", iceCream) // Input files
		// Read in the input and unmarshal to an IceCream struct

		// Compare the toppings to the toppings in toppings.json
		//Validation.Validate(ice_cream.IceCream{})
		//Validation.Validate()
		passed, err := Validation.Validate(iceCream)
		if err != nil {
			fmt.Println(err)
			return
		}
		if !passed {
			fmt.Printf("Unable to save iceCream, invalid toppings: %v", iceCream.Toppings)
			fmt.Println("\n")
		}
		fmt.Println(passed)

		// If validation fails returns false from the Validate function and return from the main function

		// If validation passes, marshall the icecream into JSON and store it in iceCreamDb.txt
		err = Repo.CreateIceCream(iceCream)
		if err != nil {
			fmt.Println(err)
			return
		}

		err = Repo.GetAllIceCreams()
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}
