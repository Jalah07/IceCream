package Repo

import (
	//"ReadFileProj/FileReader"
	"ReadFileProj/ice_cream"
	"strings"
	//"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

// Create a function that accepts an IceCream, converts it to JSON and stores it on the next available line in iceCreamDb.txt
// Return an error if any of the functions you call return an error
func CreateIceCream(ice ice_cream.IceCream) error {
	ic := ice
	output, _ := json.Marshal(ic)
	f, err := os.OpenFile("iceCreamDb.txt",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	if _, err := f.WriteString(string(output) + "\n"); err != nil {
		log.Println(err)
	}
	return err
}

// Create function that can retrieve all the ice creams from iceCreamDb.txt
// The function should unmarshal each one into an ice cream struct and return a slice of IceCream --> []IceCream
// Return an error if any of the functions you call return an error
func GetAllIceCreams()  ([]IceCream, error) {
 	file, err := ioutil.ReadFile("iceCreamDb.txt")

	if err != nil {
		log.Fatalf("Unable to read file: %v", err)
	}

	result := ice_cream.IceCream{}
	_ = json.Unmarshal([]byte(file), &result)
	fmt.Println("This is data from GetAllIceCreams", string(file))
	return err 
}

