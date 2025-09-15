package furniture

type PlyWoodBed struct {
	width float64
}

func (pwb PlyWoodBed) Price() float64 {
	return pwb.width * 400
}

type FabricSofa struct {
	length float64
}

func (f FabricSofa) Price() float64 {
	return f.length * 200
}

type PlyWoodWardrobe struct {
	height, length float64
}

func (p PlyWoodWardrobe) Price() float64 {
	return p.height * p.length * 100
}
