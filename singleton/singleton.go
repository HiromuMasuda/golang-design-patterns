package singleton

import "fmt"

type registerNote struct {
	Name string
}

var instance *registerNote

func GetInstance() *registerNote {
	if instance == nil {
		instance = &registerNote{"register_note"}
	}
	return instance
}

func Run() {
	instance1 := GetInstance()
	instance2 := GetInstance()
	fmt.Println(instance1.Name)
	fmt.Println(instance2.Name)
}
