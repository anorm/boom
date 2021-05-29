package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
)

type Entity struct {
	Type        string
	Name        string
	Description string
	Shape       string
}

func main() {
	file, err := os.Open(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	decoder := yaml.NewDecoder(file)

	entity := Entity{}
	err = decoder.Decode(&entity)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entity)

	fmt.Println("Boom!")
}
