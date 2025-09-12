package furniture

type QuanYouFactory struct {
}

func (qyf QuanYouFactory) CreateBed(width float64) Bed {
	return SolidWoodBed{width: width}
}
func (qyf QuanYouFactory) CreateSofa(length float64) Sofa {
	return LeatherSofa{length: length}
}
func (qyf QuanYouFactory) CreateWardrobe(height, length float64) Wardrobe {
	return SolidWoodWardrobe{height: height, length: length}
}
