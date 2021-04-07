package main

import (
	"./UTF16Convert"
	"fmt"
	"log"
)

func main()  {
	filename := "test_file.txt"
	final, err := UTF16Convert.Replacement(filename)
	if err != nil {
		log.Fatal(err)
		return
	}
	fmt.Println(final)

}