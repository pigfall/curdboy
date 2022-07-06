package filter

import (
	"fmt"
	"github.com/xlab/treeprint"
)

type LabelValueType int

const(
	LabelValueTypeStr LabelValueType= iota+1
	LabelValueTypeNumber 
		
)

type LabelDesc interface {
	LabelTypeId() (string, error)
	LabelValue()interface{}
}

type Expr interface {
//	OpString() string
	DebugTree(tree treeprint.Tree)
//	ToEntStoreFilter() (predicate.Store, error)
	Evaluate(labels []LabelDesc) (bool, error)
	Accept(visitor Visitor)(interface{},error)
}

type BinaryLogicalExpr struct {
	Left  Expr
	Right Expr
	Op    *Token
}

func NewBinaryLogicalExpr(left, right Expr, op *Token) Expr {
	return &BinaryLogicalExpr{
		Left:  left,
		Right: right,
		Op:    op,
	}
}

func (this *BinaryLogicalExpr) Accept(visitor Visitor)(interface{},error){
	return visitor.VisitBinaryLogicalExpr(this)
}

func (this *BinaryLogicalExpr) Evaluate(labels []LabelDesc) (bool, error) {
	leftBool, err := this.Left.Evaluate(labels)
	if err != nil {
		return false, err
	}
	rightBool, err := this.Right.Evaluate(labels)
	if err != nil {
		return false, err
	}

	switch this.Op.Tpe {
	case TokenType_KW_And:
		return leftBool && rightBool, nil
	case TokenType_KW_Or:
		return leftBool || rightBool, nil
	default:
		return false, fmt.Errorf("Undefined binary logical operator %+v", this.Op)
	}
}


func (this *BinaryLogicalExpr) OpString() string {
	return this.Op.Literal
}

func (this *BinaryLogicalExpr) DebugTree(tree treeprint.Tree) {
	subTree := tree.AddBranch(this.OpString())
	this.Left.DebugTree(subTree)
	this.Right.DebugTree(subTree)
}

type ComparisionExpr struct {
	Left  *Token
	Right *Token
	Op    *Token
}

func NewComparisionExpr(left, right, op *Token) *ComparisionExpr {
	return &ComparisionExpr{
		Left:  left,
		Right: right,
		Op:    op,
	}
}

func (this *ComparisionExpr) Accept(visitor Visitor)(interface{},error){
	return visitor.VisitComparisionExpr(this)
}


func (this *ComparisionExpr) Evaluate(labels []LabelDesc) (bool, error) {
	labelMaps := make(map[string]LabelDesc)
	for _, l := range labels {
		var labelTypeId, err = l.LabelTypeId()
		if err != nil {
			return false, err
		}
		if labelMaps[labelTypeId] != nil {
			return false, fmt.Errorf("store label %s repeated ", labelTypeId)
		}
		labelMaps[labelTypeId] = l
	}

	key := this.Left.GetValue().(string)

	switch this.Op.Tpe {
	case TokenType_KW_Gt:
		return this.numberCompare(labelMaps, func(left, right float64) bool {
			return left > right
		})
	case TokenType_KW_Ge:
		return this.numberCompare(labelMaps, func(left, right float64) bool {
			return left >= right
		})
	case TokenType_KW_Eq:
		if label := labelMaps[key]; label == nil {
			return false, nil
		} else {
			return  label.LabelValue()== this.Right.GetValue(), nil
		}

	case TokenType_KW_Ne:
		if label := labelMaps[key]; label == nil {
			return false, nil
		} else {
			return label.LabelValue() != this.Right.GetValue(), nil
		}

	case TokenType_KW_Lt:
		return this.numberCompare(labelMaps, func(left, right float64) bool {
			return left < right
		})
	case TokenType_KW_Le:
		return this.numberCompare(labelMaps, func(left, right float64) bool {
			return left <= right
		})
	default:
		return false, fmt.Errorf("Undefined comparisionExpr operator type %+v", this.Op)
	}
}


func (this *ComparisionExpr) ToString() string {
	return fmt.Sprintf("{%v %v %v}", this.Left.Literal, this.Op.Literal, this.Right.Literal)
}

func (this *ComparisionExpr) OpString() string {
	return this.Op.Literal
}
func (this *ComparisionExpr) DebugTree(tree treeprint.Tree) {
	subTree := tree.AddBranch(this.OpString())
	subTree.AddNode(this.Left.Literal)
	subTree.AddNode(this.Right.Literal)
}

type UnaryExpr struct {
	Op   *Token
	Expr Expr
}

func NewUnaryExpr(op *Token, expr Expr) Expr {
	return &UnaryExpr{
		Op:   op,
		Expr: expr,
	}
}

func (this *UnaryExpr) Accept(visitor Visitor)(interface{},error){
	return visitor.VisitUnaryExpr(this)
}

func (this *UnaryExpr) Evaluate(labels []LabelDesc) (bool, error) {
	b, err := this.Expr.Evaluate(labels)
	if err != nil {
		return false, err
	}
	return !b, nil
}


func (this *UnaryExpr) ToString() string {
	return fmt.Sprintf("%v %v", this.Op, this.Expr)
}

func (this *UnaryExpr) OpString() string {
	return this.Op.Literal
}

func (this *UnaryExpr) DebugTree(tree treeprint.Tree) {
	subTree := tree.AddBranch(this.OpString())
	this.Expr.DebugTree(subTree)
}



func (this *ComparisionExpr) numberCompare(labelMaps map[string]LabelDesc, cmp func(left, right float64) bool) (bool, error) {
	key := this.Left.GetValue().(string)
	if label := labelMaps[key]; label == nil {
		return false, nil
	} else {
		exprValue := this.Right.GetNumberValue()
		labelValue,ok := label.LabelValue().(float64)
		if !ok {
			return false,fmt.Errorf("Label value is not float64")
		}
		return cmp((exprValue),float64(labelValue)), nil
	}
}

