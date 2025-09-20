package buyticket

import (
	"fmt"
	"time"
)

type Train struct {
	from       string
	to         string
	departTime time.Time
	emptySeats int
	// D G P K T
	trainType string
}

func (t *Train) String() string {
	return fmt.Sprintf("%s %s %s %d %s", t.from, t.to, t.departTime.Format("2006-01-02 15:04:05"), t.emptySeats, t.trainType)
}

func NewTrain(from, to, depart string, emptySeats int, trainType string) (*Train, error) {
	departTime, err := time.Parse("2006-01-02 15:04:05", depart)
	if err != nil {
		return nil, err
	}
	return &Train{from: from, to: to, departTime: departTime, emptySeats: emptySeats, trainType: trainType}, nil
}

type TrainIterator interface {
	hasNext() bool
	Next() *Train
}
type DefaultTrainIterator struct {
	curIndex  int
	nextIndex int
	trains    []*Train
	isTarget  func(*Train) bool
}

func NewTrainIterator(trains []*Train, isTarget func(*Train) bool) TrainIterator {
	return &DefaultTrainIterator{curIndex: -1, nextIndex: -1, trains: trains, isTarget: isTarget}
}

func (t *DefaultTrainIterator) hasNext() bool {
	for i := max(t.curIndex+1, 0); i < len(t.trains); i++ {
		if t.isTarget(t.trains[i]) {
			t.nextIndex = i
			return true
		}
	}
	return false
}
func (t *DefaultTrainIterator) Next() *Train {
	if t.nextIndex >= 0 && t.nextIndex < len(t.trains) {
		t.curIndex = t.nextIndex
		nextTrain := t.trains[t.nextIndex]
		t.nextIndex = -1
		return nextTrain
	}
	return nil
}
