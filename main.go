package main

import "github.com/Drafteame/pkl-linter/cmd/root"

func main() {
	if err := root.GetCmd().Execute(); err != nil {
		panic(err)
	}
}
