package main

import "diff-problems/infrastructure"

func main() {
	infrastructure.Initialize()
	err := infrastructure.Router.Run()
	if err != nil {
		panic(err)
	}
}
