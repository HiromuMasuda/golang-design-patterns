package main

import (
	"./adapter"
	"./facade"
	"./singleton"
	"./strategy"
	"./template_method"
)

func main() {
	adapter.RunDelegation()
	adapter.RunInheritance()
	facade.Run()
	singleton.Run()
	strategy.Run()
	template_method.Run()
}
