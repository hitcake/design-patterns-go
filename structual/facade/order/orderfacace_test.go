package order

import (
	"encoding/json"
	"os"
	"testing"
)

func TestOrderFacade_GetOrder(t *testing.T) {
	orderFacade := &OrderFacade{
		&OrderService{},
		&UserService{},
		&ProductService{},
	}
	orderWrapper, err := orderFacade.GetOrder(1)
	if err != nil {
		t.Errorf("GetOrder() error:%v", err)
	}
	json.NewEncoder(os.Stdout).Encode(*orderWrapper)
}
