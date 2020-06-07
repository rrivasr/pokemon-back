package main

import "pokemon-back/cmd"

func main() {
	if err := cmd.StartApp(); err != nil {
		print("Error "+ err.Error())
	}
}
