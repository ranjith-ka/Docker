package main

import (
	"encoding/json"
	"fmt"
)

type Phone struct {
	Model             string
	Cost              int
	Brand             string
	Availablity_place []string
}

func main() {
	p1 := Phone{
		Model: "narzo20pro",
		Cost:  200,
		Brand: "realme",
		Availablity_place: []string{
			"bangalore",
			"karaikal",
		},
	}
	p2 := Phone{
		Model: "narzo30pro",
		Cost:  300,
		Brand: "realme",
		Availablity_place: []string{
			"bangalore",
			"karaikal",
		},
	}

	mob := []Phone{p1, p2}
	fmt.Println(mob)

	b, err := json.Marshal(mob)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(b))
}
