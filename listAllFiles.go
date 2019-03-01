package main

import (
	"fmt"
	"io/ioutil"
	"log"
)


func main() {
	files, err := ioutil.ReadDir("/Users/mcclayac/Google Drive/images")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {

		if file.IsDir() {
			//fmt.Println(file.Name() + " is Directory")
			continue
		}
		fmt.Println(file.Mode(), file.Name())
	}
}


