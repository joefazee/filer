package main

import (
	"flag"
	"fmt"
	"log"
)

const (
	ActionSplit = "split"
	ActionJoin  = "join"
)

func main() {

	var fileName string
	var action string
	var partSize int64

	flag.StringVar(&fileName, "file", "", "-file=name the name of the file to split or join")
	flag.StringVar(&action, "action", "", "-action=split action can be split or join")
	flag.Int64Var(&partSize, "size", 0, "-size the size in bytes to chunk the file")

	flag.Parse()

	if fileName == "" {
		log.Fatal("fileName is required")
	}

	if action == ActionSplit && partSize == 0 {
		log.Fatal("size is required in bytes")
	}

	var err error
	switch action {
	case ActionSplit:
		err = SplitFile(fileName, partSize)
	case ActionJoin:
		err = JoinFiles(fileName)
	default:
		log.Fatal("invalid action")
	}

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Worked")
}
