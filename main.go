package main

import "fmt"

type test struct {
	On    bool
	Ammo  int
	Power int
}

func (a *test) Shoot() bool {
	if a.On == false || a.Ammo == 0 {
		return false
	}
	a.Ammo--
	return true
}
func (a *test) RideBike() bool {
	if a.On == false || a.Power == 0 {
		return false
	}
	a.Power--
	return true
}

func main() {
	testStruct := &test{}
	fmt.Print(testStruct.Shoot(), testStruct.RideBike())
}
