package main

type ChairFactory interface {
	Set() error
}

type SofaFactory interface {
	Filling(t string) error
	Set() error
}

type TableFactory interface {
	SetHigh(h int) error
	SetLong(l int) error
	SetWide(w int) error
}

type FurnitureFactory interface {
	CreateChairFac() ChairFactory
	CreateSofaFac() SofaFactory
	CreateTableFac() TableFactory
}

func GetFurnitureFactory(t string) FurnitureFactory {
	switch t {
	case "ancient":
		return NewAncientStyleFurnitureFactory()
	case "european":
		return NewEuropeanStyleFurnitureFactory()
	default:
		return NewAncientStyleFurnitureFactory()
	}
}

//古风家具工厂
type AncientStyleFurnitureFactory struct {
}

func NewAncientStyleFurnitureFactory() *AncientStyleFurnitureFactory {
	return &AncientStyleFurnitureFactory{}
}

func (a *AncientStyleFurnitureFactory) CreateChairFac() ChairFactory {
	return nil
}
func (a *AncientStyleFurnitureFactory) CreateSofaFac() SofaFactory {
	return nil
}
func (a *AncientStyleFurnitureFactory) CreateTableFac() TableFactory {
	return nil
}

//欧式风格家具工厂
type EuropeanStyleFurnitureFactory struct {
}

func NewEuropeanStyleFurnitureFactory() *EuropeanStyleFurnitureFactory {
	return &EuropeanStyleFurnitureFactory{}
}

func (e *EuropeanStyleFurnitureFactory) CreateChairFac() ChairFactory {
	return nil
}
func (e *EuropeanStyleFurnitureFactory) CreateSofaFac() SofaFactory {
	return nil
}
func (e *EuropeanStyleFurnitureFactory) CreateTableFac() TableFactory {
	return nil
}
