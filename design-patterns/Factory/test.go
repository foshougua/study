package main

import "fmt"

type Jacket interface {
	GetType() string
	Show()
}

type Sweater struct{}

func NewSweater() *Sweater {
	return &Sweater{}
}
func (s *Sweater) GetType() string {
	return "sweater"
}

func (s *Sweater) Show() {
	fmt.Println("it is sweater ...")
}

type TShirt struct{}

func NewTShirt() *TShirt {
	return &TShirt{}
}

func (t *TShirt) GetType() string {
	return "t-shirt"
}

func (t *TShirt) Show() {
	fmt.Println("it is t-shirt ...")
}

func NewJacket(t string) Jacket {
	switch t {
	case "sweater":
		return NewSweater()
	case "t-shirt":
		return NewTShirt()
	}
	return nil
}
