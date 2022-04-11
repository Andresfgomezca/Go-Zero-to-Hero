package main

import (
	"errors"
	"fmt"
	"go-exercises/files"
	"os"
)

func checkIfCommandIsValid(arguments ...string) error {
	if len(arguments) < 2 {
		return errors.New("COMMAND MUST BE WITH TWO ARGUMENTS OR MORE")
	}
	return nil
}

func main() {
	var f files.Files

	Arguments := os.Args[1:]
	e := checkIfCommandIsValid(Arguments...)

	if e != nil {
		fmt.Println(e)
	}

	f = files.Files{FilePath: Arguments[0]}

	switch s := Arguments[1]; s {
	case "create":
		f.CreateFile()
	case "edit":
		f.EditFile(Arguments[2], Arguments[3])
	case "rename":
		f.RenameFile(Arguments[2])
	case "read":
		fmt.Println(string(f.ReadFile()))
	case "remove":
		f.RemoveFile()
	default:
		fmt.Println("Unrecognized command")
	}

}
