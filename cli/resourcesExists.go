package main

import (
	"fmt"
	"log"
	"os"
)

func (cfg *config) resourcesExists() {
	fmt.Println("Resources already exists!!!")
	fmt.Printf("DELETE current resources and Create NEW??(y/N):")
	var opt string
	fmt.Scanln(&opt)
	if opt != "y" {
		fmt.Println("Exiting...")
	} else {
		fmt.Println("Deleting resources...")
		err := os.Remove(cfg.dbName)
		if err != nil {
			log.Fatal(err)
		}
		cfg.createResources()
	}
}
