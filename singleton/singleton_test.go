package singleton

import "testing"

func TestSingleton(t *testing.T) {
	instance1 := GetInstance()
	instance2 := GetInstance()

	result := instance1.Name
	expect := instance2.Name

	if result != expect {
		t.Errorf("Expect result to equal %s, but %s.", expect, result)
	}
}
