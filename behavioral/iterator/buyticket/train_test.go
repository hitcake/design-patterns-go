package buyticket

import (
	"fmt"
	"testing"
)

func TestIterator(t *testing.T) {
	g101, err := NewTrain("北京南", "上海红桥", "2025-10-01 06:10:00", 0, "G")
	if err != nil {
		t.Error(err)
	}
	g103, err := NewTrain("北京南", "上海红桥", "2025-10-01 06:20:00", 50, "G")
	if err != nil {
		t.Error(err)
	}
	g1, err := NewTrain("北京南", "上海", "2025-10-01 07:00:00", 100, "G")
	if err != nil {
		t.Error(err)
	}
	g3, err := NewTrain("北京", "上海", "2025-10-01 08:00:00", 100, "G")
	if err != nil {
		t.Error(err)
	}
	trains := []*Train{g101, g103, g1, g3}

	hasSeatFilter := func(t *Train) bool {
		return t.emptySeats > 0
	}
	assertTrue(t, NewTrainIterator(trains, hasSeatFilter), 3)
	departFromBJFilter := func(t *Train) bool {
		return t.from == "北京"
	}
	assertTrue(t, NewTrainIterator(trains, departFromBJFilter), 1)
}

func assertTrue(t *testing.T, iterator TrainIterator, expected int) {
	count := 0
	for iterator.hasNext() {
		count++
		t := iterator.Next()
		fmt.Println(t)
	}
	if count != expected {
		t.Errorf("count %d expected %d", count, expected)
	}
}
