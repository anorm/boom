package main

import (
	"fmt"
	"log"
	"os"
	"strings"

    "github.com/anorm/boom/v2/pkg/workspace"

	"gopkg.in/yaml.v3"
)

type Workspace struct {
    Name string
    Description string

    Landscape Landscape
    Perspectives []Perspective
}

type Landscape struct {
	Name        string
	Description string
}

type Perspective struct {
	Name        string
	Description string
}

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
}

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func parse(filename string) (<-chan Entity, error) {
    ch := make(chan Entity)
    go func() {
        ch<-Entity{}
        close(ch)
    }()
    return ch, nil
}

func main() {
    workspaceFilename := os.Args[1]

    ws, err := workspace.Load(workspaceFilename)
	if err != nil {
		log.Fatal(err)
	}
    fmt.Println(ws)

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
