package furniture

type SolidWoodBed struct {
	width float64
}

func (swb SolidWoodBed) Price() float64 {
	return swb.width * 1800
}

type LeatherSofa struct {
	length float64
}

func (ls LeatherSofa) Price() float64 {
	return ls.length * 1500
}

type SolidWoodWardrobe struct {
	height, length float64
}

func (sww SolidWoodWardrobe) Price() float64 {
	return sww.height * sww.length * 600
}
