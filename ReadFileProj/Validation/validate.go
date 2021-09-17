package Validation

import (
	"ReadFileProj/ice_cream"
	"encoding/json"
	//"fmt"
	"io/ioutil"
	"log"
	"strings"
	//"ReadFileProj/FileReader"
	//"main.go/FileReader"
	//"strings"
)

// Write a function that takes in an ice cream then reads the file toppings.json.
// Iterate through the toppings returned from toppings.json and iterate throught the toppings in the IceCream
// to verify that the IceCream ony contains toppings listed in toppings.json.
// If the ice cream fails validation return false, otherwise return true.
func Validate(iceCream ice_cream.IceCream) (bool, error) {
	// Read file toppings.json
	file, err := ioutil.ReadFile("input/toppings.json")

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	data := ice_cream.IceCream{}
	_ = json.Unmarshal([]byte(file), &data)

	

	for _, v := range iceCream.Toppings {
		if contains(data.Toppings, v) {
			continue
		}else if !contains(data.Toppings, v) {
			return false, err
		}
		return true, err
	}
	return true, err
}

// Checks whether the slice contains the string
func contains(s []string, e string) bool {
	for _, a := range s {
		if strings.Contains(a, e){
			return true

		}
	}
	return false

}
