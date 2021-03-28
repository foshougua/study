package main

type HouseBuilder interface {
	SetWall(w string) error
	SetGarden(t string) error
	Renovation(t string) error
	GetHouse() *House
}

type Director interface {
	GetHouse(label string) *House
}

//公寓构造器
type ApartmentHouse struct {
	Wall           string
	GardenType     string
	RenovationType string
}

func (a *ApartmentHouse) SetWall(w string) error {
	a.Wall = w
	return nil
}

func (a *ApartmentHouse) SetGarden(t string) error {
	a.GardenType = t
	return nil
}

func (a *ApartmentHouse) Renovation(t string) error {
	a.RenovationType = t
	return nil
}

func (a *ApartmentHouse) GetHouse() *House {
	return NewHouse(a.Wall, a.GardenType, a.RenovationType)
}

//别墅
type VillaHouse struct {
	Wall           string
	GardenType     string
	RenovationType string
}

func (v *VillaHouse) SetWall(w string) error {
	v.Wall = w
	return nil
}

func (v *VillaHouse) SetGarden(t string) error {
	v.GardenType = t
	return nil
}

func (v *VillaHouse) Renovation(t string) error {
	v.RenovationType = t
	return nil
}

func (v *VillaHouse) GetHouse() *House {
	return NewHouse(v.Wall, v.GardenType, v.RenovationType)
}

//管理员
type HouseDirector struct {
	b HouseBuilder
}

func NewHouseDirector(b HouseBuilder) *HouseDirector {
	return &HouseDirector{b: b}
}

func (h *HouseDirector) GetHouse(label string) *House {
	_ = h.b.SetWall(label + " wall")
	_ = h.b.SetGarden(label + " garden")
	_ = h.b.Renovation(label + " renovation")
	return h.b.GetHouse()
}

type House struct {
	Wall           string
	GardenType     string
	RenovationType string
}

func NewHouse(w, gt, rt string) *House {
	return &House{Wall: w, GardenType: gt, RenovationType: rt}
}
