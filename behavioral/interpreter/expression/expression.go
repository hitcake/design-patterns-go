package expression

type Expression interface {
	Interpret() int
}
type NumberExpression struct {
	number int
}

func (n *NumberExpression) Interpret() int {
	return n.number
}

type AddExpression struct {
	left, right Expression
}

func (a *AddExpression) Interpret() int {
	return a.left.Interpret() + a.right.Interpret()
}

type SubtractExpression struct {
	left, right Expression
}

func (s *SubtractExpression) Interpret() int {
	return s.left.Interpret() - s.right.Interpret()
}
