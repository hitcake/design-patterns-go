package furniture

import (
	"fmt"
	"testing"
)

func TestByFurniture(t *testing.T) {
	fmt.Println("进入全友家居")
	factory := QuanYouFactory{}
	bed := factory.CreateBed(1.8)
	fmt.Println(bed.Price())
	sofa := factory.CreateSofa(3.2)
	fmt.Println(sofa.Price())
	wardrobe := factory.CreateWardrobe(2.5, 1.6)
	fmt.Println(wardrobe.Price())
	fmt.Println("进入兔宝宝家居")
	factory2 := TuBaoFactory{}
	bed = factory2.CreateBed(1.8)
	fmt.Println(bed.Price())
	sofa = factory2.CreateSofa(3.2)
	fmt.Println(sofa.Price())
	wardrobe = factory2.CreateWardrobe(2.5, 1.6)
	fmt.Println(wardrobe.Price())
}
