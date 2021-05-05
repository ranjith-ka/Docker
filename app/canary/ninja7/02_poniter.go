package main

import "fmt"

type realme struct {
	model string
}

func main() {
	m1 := realme{
		model: "narzo30pro",
	}
	fmt.Println(m1)
	s := Changeme(&m1)
	fmt.Println(s)
}

// Changeme method takes the pointer values and change the underlying address value
func Changeme(p *realme) string {
	p.model = "narzo20"
	s1 := p.model
	return s1
	// (*p).model = "relame7"
}
