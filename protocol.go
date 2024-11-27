package main

import "fmt"

type Command struct {
}

func parseCommand(msg string) (Command, error) {
	t := msg[0]
	fmt.Println(t)
	return Command{}, nil
}
