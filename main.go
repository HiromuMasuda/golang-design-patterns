package main

import (
	"./adapter"
	"./facade"
	"./strategy"
	"./template_method"
)

func main() {
	adapter.RunDelegation()
	adapter.RunInheritance()
	facade.Run()
	strategy.Run()
	template_method.Run()
}
