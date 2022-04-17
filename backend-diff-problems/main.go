package main

import "diff-problems/infrastructure"

func main() {
	err := infrastructure.Router.Run()
	if err != nil {
		panic(err)
	}
}
