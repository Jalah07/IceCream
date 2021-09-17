package FileReader

import (
	"encoding/json"
	"io/ioutil"
	"log"
	//"ReadFile/ice_cream"
	"ReadFileProj/ice_cream"
	//"testing/quick"
	//"fmt"

)

//Use the appropriate functions from the os package, JSON package, ioutils package etc. To read the input
// from a file, unmarshal it into an IceCream struct and return the struct
// Return an error if any of your functions that you call return an error
func ReadFile(fileName string)  (ice_cream.IceCream, error) {
	file, err := ioutil.ReadFile(fileName)

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	data := ice_cream.IceCream{}
	_ = json.Unmarshal([]byte(file), &data)
	//fmt.Println("This is data from input.go",data)
	//fmt.Println("This is from ReadFile toppings", data.Toppings)
	return data, err
}
