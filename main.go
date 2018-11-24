package main

import (
	"./adapter"
	"./facade"
	"./strategy"
)

func main() {
	adapter.RunDelegation()
	adapter.RunInheritance()
	facade.Run()
	strategy.Run()
}
