package main

import "fmt"

func main() {
	fmt.Println("test abstract factory ...")
	fac := GetFurnitureFactory("ancient")
	fmt.Println(fac.CreateChairFac())
	fmt.Println(fac.CreateSofaFac())
	fmt.Println(fac.CreateTableFac())
}
