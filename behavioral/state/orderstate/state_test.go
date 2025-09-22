package orderstate

import "testing"

func TestState(t *testing.T) {
	order := NewOrder("123123123")
	err := order.Pay()
	if err != nil {
		t.Error(err)
	}
	err = order.Complete()
	if err == nil {
		t.Error("paid should not complete")
	}
	err = order.Ship()
	if err != nil {
		t.Error(err)
	}
	err = order.Complete()
	if err != nil {
		t.Error(err)
	}
}
