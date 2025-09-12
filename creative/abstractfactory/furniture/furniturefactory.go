package furniture

type FurnitureFactory interface {
	CreateBed(width float64) Bed
	CreateSofa(length float64) Sofa
	CreateWardrobe(height, length float64) Wardrobe
}
