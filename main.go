package main

import "github.com/sudhanva-nadiger/task/cmd"

func main() {
	err := cmd.RootCmd.Execute()

	if err != nil {
		panic(err)
	}
}
