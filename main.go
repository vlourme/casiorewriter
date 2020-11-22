package main

import (
	"casiorewriter/cmd"
	"fmt"
	"github.com/jaffee/commandeer"
)

func main() {
	err := commandeer.Run(cmd.NewMain())
	if err != nil {
		fmt.Println(err)
	}
}