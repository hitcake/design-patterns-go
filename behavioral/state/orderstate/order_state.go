package orderstate

import "errors"

type OrderState interface {
	Pay(order *Order) error
	Cancel(order *Order) error
	Ship(order *Order) error
	Complete(order *Order) error
}
type PendingPaymentState struct{}

func (s *PendingPaymentState) Pay(order *Order) error {
	order.setState(&PaidState{})
	return nil
}
func (s *PendingPaymentState) Cancel(order *Order) error {
	order.setState(&CanceledState{})
	return nil
}
func (s *PendingPaymentState) Ship(order *Order) error {
	return errors.New("no paid , can't ship")
}
func (s *PendingPaymentState) Complete(order *Order) error {
	return errors.New("no paid , cannt' complete")
}

type PaidState struct{}

func (s *PaidState) Pay(order *Order) error {
	return errors.New("already paid , can't pay again")
}
func (s *PaidState) Cancel(order *Order) error {
	order.setState(&CanceledState{})
	return nil
}
func (s *PaidState) Ship(order *Order) error {
	order.setState(&ShippedState{})
	return nil
}
func (s *PaidState) Complete(order *Order) error {
	return errors.New("no ship , can't complete")
}

type CanceledState struct{}

func (s *CanceledState) Pay(order *Order) error {
	return errors.New("already canceled , can't pay again")
}
func (s *CanceledState) Cancel(order *Order) error {
	return errors.New("already canceled , can't cancel")
}
func (s *CanceledState) Ship(order *Order) error {
	return errors.New("already canceled , can't ship")
}
func (s *CanceledState) Complete(order *Order) error {
	return errors.New("already canceled , can't complete")
}

type ShippedState struct{}

func (s *ShippedState) Pay(order *Order) error {
	return errors.New("already shipped , can't pay again")
}
func (s *ShippedState) Cancel(order *Order) error {
	return errors.New("already shipped , can't cancel")
}
func (s *ShippedState) Ship(order *Order) error {
	return errors.New("already shipped , can't ship")
}
func (s *ShippedState) Complete(order *Order) error {
	order.setState(&CompleteState{})
	return nil
}

type CompleteState struct{}

func (s *CompleteState) Pay(order *Order) error {
	return errors.New("already paid , can't pay again")
}
func (s *CompleteState) Cancel(order *Order) error {
	return errors.New("already shipped , can't cancel")
}
func (s *CompleteState) Ship(order *Order) error {
	return errors.New("already shipped , can't ship")
}
func (s *CompleteState) Complete(order *Order) error {
	return errors.New("already complete , can't complete")
}
