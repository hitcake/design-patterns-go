package furniture

type TuBaoFactory struct {
}

func (f TuBaoFactory) CreateBed(width float64) Bed {
	return PlyWoodBed{width: width}
}

func (f TuBaoFactory) CreateSofa(length float64) Sofa {
	return FabricSofa{length: length}
}

func (f TuBaoFactory) CreateWardrobe(height, length float64) Wardrobe {
	return PlyWoodWardrobe{height: height, length: length}
}
