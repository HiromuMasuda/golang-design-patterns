package main

import (
	"./adapter"
	"./facade"
)

func main() {
	adapter.RunDelegation()
	adapter.RunInheritance()
	facade.Run()
}
