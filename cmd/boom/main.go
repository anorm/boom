package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/anorm/boom/v2/pkg/model"
	"github.com/anorm/boom/v2/pkg/workspace"
	"gopkg.in/yaml.v3"
)

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}

func parse(filename string) (<-chan model.Entity, error) {
	ch := make(chan model.Entity)
	go func() {
		ch <- model.Entity{}
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

	entity := model.Entity{}
	err = decoder.Decode(&entity)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(entity)

	fmt.Println("Boom!")
}
