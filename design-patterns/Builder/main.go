package main

func main() {
	b := NewHouseDirector(&ApartmentHouse{})
	b.GetHouse("apartment")

	b1 := NewHouseDirector(&VillaHouse{})
	b1.GetHouse("villa")

	return
}
