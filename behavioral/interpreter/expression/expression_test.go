package expression

import "testing"

func TestExpr(t *testing.T) {
	five := &NumberExpression{number: 5}
	ten := &NumberExpression{number: 10}
	fifteen := &NumberExpression{number: 15}

	expr := SubtractExpression{left: &AddExpression{five, ten}, right: fifteen}
	result := expr.Interpret()
	if result != 0 {
		t.Errorf("Expression interpretation failed. Expected: %d, Actual: %d", 0, result)
	}
}
