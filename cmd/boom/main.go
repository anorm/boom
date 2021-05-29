package main

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
    "strings"
)

type Entity struct {
	Type        string
	Name        string
	Description string
	Tech        string
	Shape       string
	Relations   []struct {
		Source      string
		Target      string
		Description string
		Tech        string
	}
    children []Entity
}

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
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
