package filter

type Visitor interface {
	VisitBinaryLogicalExpr(*BinaryLogicalExpr) (interface{}, error)
	VisitComparisionExpr(*ComparisionExpr) (interface{}, error)
	VisitUnaryExpr(*UnaryExpr) (interface{}, error)
}
