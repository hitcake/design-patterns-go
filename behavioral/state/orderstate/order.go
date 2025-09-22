package orderstate

type Order struct {
	orderId string
	state   OrderState
}

func (o *Order) setState(state OrderState) {
	o.state = state
}
func NewOrder(orderId string) *Order {
	return &Order{orderId: orderId, state: &PendingPaymentState{}}
}
func (o *Order) Pay() error {
	return o.state.Pay(o)
}
func (o *Order) Cancel() error {
	return o.state.Cancel(o)
}
func (o *Order) Ship() error {
	return o.state.Ship(o)
}
func (o *Order) Complete() error {
	return o.state.Complete(o)
}
